package webtransport

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/quic-go/qpack"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/quicvarint"
)

func sendSettingsFrame(sess quic.Connection) {
	// server should send HTTP SETTINGS frame to client
	log.Debug("[1] Send SETTINGS frame")
	// https://www.ietf.org/archive/id/draft-ietf-webtrans-http3-02.html#section-3.1
	// In order to indicate support for WebTransport, both the client and the server MUST send a SETTINGS_ENABLE_WEBTRANSPORT value set to "1" in their SETTINGS frame.
	// [1] send SETTINGS frame
	// https://www.w3.org/TR/webtransport/#webtransport-constructor
	// 6. Wait for connection to receive the first SETTINGS frame, and let settings be a dictionary that represents the SETTINGS frame.
	// 7. If settings doesn’t contain SETTINGS_ENABLE_WEBTRANPORT with a value of 1, or it doesn’t contain H3_DATAGRAM with a value of 1, then abort the remaining steps and queue a network task with transport to run these steps:
	respStream, err := sess.OpenUniStream()
	if err != nil {
		log.Error("sess.OpenUniStream error: %v", err)
		return
	}

	// https://datatracker.ietf.org/doc/draft-ietf-masque-h3-datagram/
	// Implementations of HTTP/3 that support HTTP Datagrams can indicate
	// that to their peer by sending the H3_DATAGRAM SETTINGS parameter with
	// a value of 1.  The value of the H3_DATAGRAM SETTINGS parameter MUST
	// be either 0 or 1.  A value of 0 indicates that HTTP Datagrams are not
	// supported.  An endpoint that receives the H3_DATAGRAM SETTINGS
	// parameter with a value that is neither 0 or 1 MUST terminate the
	// connection with error H3_SETTINGS_ERROR.
	// buf := &bytes.Buffer{}
	buf := make([]byte, 0, 64)
	// https://quicwg.org/base-drafts/draft-ietf-quic-http.html#section-6.2.1
	// A control stream is indicated by a stream type of 0x00. Data on this stream consists of HTTP/3 frames, as defined in Section 7.2.
	// Each side MUST initiate a single control stream at the beginning of the connection and send its SETTINGS frame as the first frame on this stream. If the first frame of the control stream is any other frame type, this MUST be treated as a connection error of type H3_MISSING_SETTINGS. Only one control stream per peer is permitted; receipt of a second stream claiming to be a control stream MUST be treated as a connection error of type
	buf = quicvarint.Append(buf, 0x00)
	// https://quicwg.org/base-drafts/draft-ietf-quic-http.html#name-http-framing-layer
	// 7. HTTP Framing Layer
	// HTTP/3 Frame Format {
	// 	 Type (i),
	// 	 Length (i),
	// 	 Frame Payload (..),
	// }
	// https://quicwg.org/base-drafts/draft-ietf-quic-http.html#name-settings
	// 7.2.4. SETTINGS
	// The SETTINGS frame (type=0x04) conveys configuration parameters that affect how endpoints communicate, such as preferences and constraints on peer behavior. Individually, a SETTINGS parameter can also be referred to as a "setting"; the identifier and value of each setting parameter can be referred to as a "setting identifier" and a "setting value".
	buf = quicvarint.Append(buf, 0x04)
	var l uint64
	// H3_DATAGRAM
	// https://datatracker.ietf.org/doc/html/draft-ietf-masque-h3-datagram-05#section-9.1
	// +==============+==========+===============+=========+
	// | Setting Name | Value    | Specification | Default |
	// +==============+==========+===============+=========+
	// | H3_DATAGRAM  | 0xffd277 | This Document | 0       |
	// +--------------+----------+---------------+---------+
	l += uint64(quicvarint.Len(0xffd277) + quicvarint.Len(1))
	// SETTINGS_ENABLE_WEBTRANPORT
	// https://www.ietf.org/archive/id/draft-ietf-webtrans-http3-02.html#section-8.2
	// The SETTINGS_ENABLE_WEBTRANSPORT parameter indicates that the specified HTTP/3 connection is
	// WebTransport-capable.
	// Setting Name:ENABLE_WEBTRANSPORT
	// Value:0x2b603742
	// Default:0
	l += uint64(quicvarint.Len(0x2b603742) + quicvarint.Len(1))
	// // ???
	// l += uint64(quicvarint.Len(0x276) + quicvarint.Len(1))
	// write Length
	buf = quicvarint.Append(buf, l)
	// Write value
	// https://quicwg.org/base-drafts/draft-ietf-quic-http.html#name-settings
	// The payload of a SETTINGS frame consists of zero or more parameters. Each parameter consists of a setting identifier and a value, both encoded as QUIC variable-length integers.
	//
	// Setting {
	//   Identifier (i),
	//   Value (i),
	// }

	// SETTINGS Frame {
	//   Type (i) = 0x04,
	//   Length (i),
	//   Setting (..) ...,
	// }
	//
	// quicvarint.Write(buf, 0x276)
	// quicvarint.Write(buf, 1)
	buf = quicvarint.Append(buf, 0xffd277) // H3_DATAGRAM
	buf = quicvarint.Append(buf, 1)
	buf = quicvarint.Append(buf, 0x2b603742) // SETTINGS_ENABLE_WEBTRANSPORT
	buf = quicvarint.Append(buf, 1)

	log.Debug("\t[len=%d] %# x", len(buf), buf)
	_, err = respStream.Write(buf)
	if err != nil {
		log.Error("sendSettingsFrame error: %v", err)
	}
	log.Debug("\tSettings frame sent")
}

func receiveSettingsFrame(sess quic.Connection) error {
	recvSettingStream, _ := sess.AcceptUniStream(sess.Context())
	log.Debug("[2] receive client SETTINGS frame")
	sqr := quicvarint.NewReader(recvSettingStream)
	// stream type should = 0x00, control stream
	sty, err := quicvarint.Read(sqr)
	if err != nil {
		return err
	}
	log.Debug("\tStreamType: %# x\r", sty)
	// frame type should = 0x04, SETTINGS frame
	ftype, err := quicvarint.Read(sqr)
	if err != nil {
		return err
	}
	log.Debug("\tFrameType: %# x\r", ftype)
	// Settings length
	flen, err := quicvarint.Read(sqr)
	if err != nil {
		return err
	}
	log.Debug("\tLength: %# x(oct=%d)\r", flen, flen)
	// Frame Payload ...
	// total length is `flen`
	settingsPayload := make(map[uint64]uint64)
	payloadBuf := make([]byte, flen)
	if _, err := io.ReadFull(recvSettingStream, payloadBuf); err != nil {
		return err
	}
	bb := bytes.NewReader(payloadBuf)
	for bb.Len() > 0 {
		identifier, err := quicvarint.Read(bb)
		if err != nil {
			return err
		}
		value, err := quicvarint.Read(bb)
		if err != nil {
			return err
		}
		settingsPayload[identifier] = value
		log.Debug("\tidentifier:%# x, value: %d (%#x)\r", identifier, value, value)
	}

	return nil
}

func readHeaderFrame(reqStream io.Reader) ([]qpack.HeaderField, error) {
	// https://quicwg.org/base-drafts/draft-ietf-quic-http.html#name-http-framing-layer
	// HEADERS Frame {
	// 	Type (i) = 0x01,
	// 	Length (i),
	// 	Encoded Field Section (..),
	// }
	qr := quicvarint.NewReader(reqStream)
	// read header frame
	hdr, err := quicvarint.Read(qr)
	if err != nil {
		return nil, err
	}
	log.Debug("\theader: %# x", hdr)
	// read the length of the header block
	headerBlockLength, err := quicvarint.Read(qr)
	if err != nil {
		log.Error("readHeaderFrame error: %v", err)
	}
	log.Debug("\theader block: %# x", headerBlockLength)

	// header frame id is 0x01
	if hdr != 0x01 {
		return nil, errors.New("not header frame, should force close connection")
	}

	headerBlock := make([]byte, headerBlockLength)
	if _, err = io.ReadFull(reqStream, headerBlock); err != nil {
		return nil, err
	}
	decoder := qpack.NewDecoder(nil)
	return decoder.DecodeFull(headerBlock)
}

func writeResponseHeaderFrame(w io.Writer, status int) error {
	// https://www.ietf.org/archive/id/draft-ietf-webtrans-http3-02.html#name-negotiating-the-draft-versi
	// The header corresponding to the
	// version described in this draft is Sec-Webtransport-Http3-Draft02;
	// its value SHALL be 1. The server SHALL reply with a Sec-
	// Webtransport-Http3-Draft header indicating the selected version; its
	// value SHALL be draft02 for the version described in this draft.
	respHeader := http.Header{}
	respHeader.Add("Sec-Webtransport-Http3-Draft", "draft02")

	// From the client's perspective, a WebTransport session is established
	// when the client receives a 2xx response.  From the server's
	// perspective, a session is established once it sends a 2xx response.
	var qpackHeaders bytes.Buffer
	encoder := qpack.NewEncoder(&qpackHeaders)
	encoder.WriteField(qpack.HeaderField{
		Name:  ":status",
		Value: strconv.Itoa(status),
	})
	for k, v := range respHeader {
		for index := range v {
			encoder.WriteField(qpack.HeaderField{
				Name:  strings.ToLower(k),
				Value: v[index],
			})
		}
	}

	// buf := &bytes.Buffer{}
	buf := make([]byte, 0, 64)
	// https://www.rfc-editor.org/rfc/rfc9114.html#section-7.2.2
	// HEADERS Frame {
	// 	Type (i) = 0x01,
	// 	Length (i),
	// 	Encoded Field Section (..),
	// }
	buf = quicvarint.Append(buf, 0x01)
	buf = quicvarint.Append(buf, uint64(qpackHeaders.Len()))

	respWriter := bufio.NewWriter(w)
	if _, err := respWriter.Write(buf); err != nil {
		return err
	}
	if _, err := respWriter.Write(qpackHeaders.Bytes()); err != nil {
		return err
	}
	if err := respWriter.Flush(); err != nil {
		return err
	}

	log.Debug("[4] Response HEADER frame with status:%d", status)
	log.Debug("\t%v", respHeader)

	return nil
}
