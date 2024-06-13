// Package webtransport runs the webtrans server service
package webtransport

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"net/url"
	"time"

	"github.com/quic-go/quic-go"

	"yomo.run/prscd/chirp"
	"yomo.run/prscd/util"
)

var log = util.Log

// ListenAndServe create webtransport server
func ListenAndServe(addr string, tlsConfig *tls.Config) {
	quicConfig := &quic.Config{
		EnableDatagrams:    true,
		KeepAlivePeriod:    30 * time.Second,
		MaxIncomingStreams: 10000,
		MaxIdleTimeout:     6 * time.Second, // when Read timeout
	}

	ln, err := quic.ListenAddr(addr, tlsConfig, quicConfig)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Info("Prscd - WebTransport Server - Listening on %s", ln.Addr())
	log.Debug("tls: %+v", tlsConfig.NextProtos)

	// processing request
	for {
		sess, err := ln.Accept(context.Background())
		if err != nil {
			log.Error("ln.accept error: %s", err)
			continue
		}
		log.Info("+Session: %s", sess.RemoteAddr().String())
		go handleConnection(sess)
	}
}

func handleConnection(sess quic.Connection) {
	closeReason := "cc-88-cc"
	defer sess.CloseWithError(0x2, closeReason)
	log.Debug("handleConnection: %s", sess.RemoteAddr().String())
	// https://www.ietf.org/archive/id/draft-ietf-webtrans-http3-02.html#section-2
	// 2. Protocol Overview
	// When an HTTP/3 connection is established, both the client and server have
	// to send a SETTINGS_ENABLE_WEBTRANSPORT setting in order to indicate that
	// they both support WebTransport over HTTP/3.

	// WebTransport sessions are initiated inside a given HTTP/3 connection by
	// the client, who sends an extended CONNECT request [RFC8441]. If the server
	// accepts the request, an WebTransport session is established.
	// all the first is create a http3 connection:

	// https://quicwg.org/base-drafts/draft-ietf-quic-http.html#section-6
	// In version 1 of QUIC, the stream data containing HTTP frames is carried
	// by QUIC STREAM frames, but this framing is invisible to the HTTP framing layer.

	// right now, a QUIC connection has been established. So we focus on HTTP frames.

	// https://quicwg.org/base-drafts/draft-ietf-quic-http.html#section-3.2
	// While connection-level options pertaining to the core QUIC protocol are
	// set in the initial crypto handshake, HTTP/3-specific settings are conveyed
	// in the SETTINGS frame. After the QUIC connection is established, a SETTINGS
	// frame (Section 7.2.4) MUST be sent by each endpoint as the initial frame
	// of their respective HTTP control stream; see Section 6.2.1.

	// so,
	// Step 1: Server send SETTINGS frame
	go sendSettingsFrame(sess)

	// Step 2: Server receive SETTINGS frame from client
	err := receiveSettingsFrame(sess)
	if err != nil {
		log.Error("webtrans|receiveSettingsFrame error: %s", err)
		closeReason = "error in receive settings frame"
		return
	}

	// Step 3: wait for reading client HTTP CONNECT (client indicatation)
	stream, err := sess.AcceptStream(context.Background())
	if err != nil {
		log.Error("webtrans|acceptStream error: %s", err)
		closeReason = "error in accept stream"
		return
	}
	log.Debug("\trequest stream accepted: %d", stream.StreamID())

	var publicKey, userID string
	status, err := receiveHTTPConnectHeaderFrame(stream, &publicKey, &userID)
	if err != nil {
		log.Error("webtrans|receiveHTTPConnectHeaderFrame error: %s", err)
		closeReason = "error in receive http connect header frame"
		return
	}

	appID, ok := chirp.Node.AuthUser(publicKey)
	if !ok {
		status = 401
	}

	// Step 4: response HEADER frame if client is valid
	err = writeResponseHeaderFrame(stream, status)
	if err != nil {
		log.Error("webtrans|writeResponseHeaderFrame error: %s", err)
		closeReason = "error in write response header frame"
		return
	}

	if status >= 300 {
		closeReason = "failed with response status other than 2xx"
		return
	}

	log.Debug("Prepared! Start to work ... uid: %s", userID)

	// Step 5: start to processing presencejs protocol
	pconn := chirp.NewWebTransportConnection(sess)
	peer := chirp.Node.AddPeer(pconn, userID, appID)
	log.Info("[%s-%s] Upgrade done!", peer.Sid, peer.Cid)

	// Handle Datagram
	go func() {
		for {
			msg, err := sess.ReceiveMessage(context.Background())
			if err != nil {
				// ignore errors here, we will handle client close event in stream loop
				return
			}
			log.Debug("ReceiveMessage: %s", msg)
			// log.Debug("ReceiveMessage: %# x", msg)
			// be careful, the first byte of msg is 0x00
			reader := bytes.NewReader(msg[1:])
			peer.HandleSignal(reader)
		}
	}()

	for {
		var buf = make([]byte, 1024)
		_, err := stream.Read(buf)
		if err != nil {
			// if client close the connection, error will occurs here
			// for example: err:timeout: no recent network activity
			log.Error("[%s] stream.Read error: %s", pconn.RemoteAddr(), err)
			peer.Disconnect()
			stream.Close()
			sess.CloseWithError(0, "client disconnected")
			break
		}
	}
}

// [3]: wait for reading client HTTP CONNECT (client indicatation)
// https://datatracker.ietf.org/doc/html/draft-ietf-webtrans-http3/#section-3.3
// In order to create a new WebTransport session, a client can send an
// HTTP CONNECT request.  The :protocol pseudo-header field ([RFC8441])
// MUST be set to webtransport.  The :scheme field MUST be https.  Both
// the :authority and the :path value MUST be set; those fields indicate
// the desired WebTransport server. An Origin header [RFC6454] MUST be
// provided within the request.

// Upon receiving an extended CONNECT request with a :protocol field set
// to webtransport, the HTTP/3 server can check if it has a WebTransport
// server associated with the specified :authority and :path values. If
// it does not, it SHOULD reply with status code 404 (Section 6.5.4,
// [RFC7231]). If it does, it MAY accept the session by replying with a
// 2xx series status code, as defined in Section 15.3 of [SEMANTICS].
// The WebTransport server MUST verify the Origin header to ensure that
// the specified origin is allowed to access the server in question.
//
// From the client's perspective, a WebTransport session is established
// when the client receives a 2xx response. From the server's
// perspective, a session is established once it sends a 2xx response.
// WebTransport over HTTP/3 does not support 0-RTT.
func receiveHTTPConnectHeaderFrame(reqStream quic.Stream, publicKey, userID *string) (status int, err error) {
	log.Debug("[3] Receive HTTP CONNECT from client")

	// read header frame which client requested
	headers, err := readHeaderFrame(reqStream)
	if err != nil {
		return 401, err
	}

	// if developers need validate request header, below is the best place to do it
	// The :protocol pseudo-header field ([RFC8441])
	// MUST be set to webtransport.
	// The :scheme field MUST be https.
	// Both the :authority and the :path value MUST be set; those fields indicate
	// the desired WebTransport server. An Origin header [RFC6454] MUST be
	// provided within the request.
	// Which looks like:
	// 2022/02/07 11:24:59 	[header] 0: {:scheme https}
	// 2022/02/07 11:24:59 	[header] 1: {:method CONNECT}
	// 2022/02/07 11:24:59 	[header] 2: {:authority lo.yomo.dev:4433}
	// 2022/02/07 11:24:59 	[header] 3: {:path /counter}
	// 2022/02/07 11:24:59 	[header] 4: {:protocol webtransport}
	// 2022/02/07 11:24:59 	[header] 5: {sec-webtransport-http3-draft02 1}
	// 2022/02/07 11:24:59 	[header] 6: {origin https://webtransport-client.vercel.app}

	var authority, path, scheme, protocol, origin, method, version string
	for key, val := range headers {
		log.Debug("\t[header] %d: %s", key, val)
		if val.Name == ":authority" { // like prscd.yomo.dev:443
			authority = val.Value
		} else if val.Name == ":path" { // `/v1/webtrans?publickey=123&id=yomo-1`
			path = val.Value
		} else if val.Name == ":scheme" { // must be https
			scheme = val.Value
		} else if val.Name == ":method" { // CONNECT
			method = val.Value
		} else if val.Name == ":protocol" { // must be webtransport
			protocol = val.Value
		} else if val.Name == "origin" { // origin of client
			origin = val.Value
		} else if val.Name == "sec-webtransport-http3-draft02" { // must be 1
			version = val.Value
		}
	}

	if protocol != "webtransport" {
		return 401, errors.New("protocol has to be webtransport")
	}

	if scheme != "https" {
		return 401, errors.New("scheme has to be https")
	}

	if method != "CONNECT" {
		return 401, errors.New("method has to be CONNECT")
	}

	if version != "1" {
		return 401, errors.New("sec-webtransport-http3-draft02 has to be 1")
	}

	// if origin need to be validated, do it here
	log.Debug("origin: %s", origin)

	// by checking authority, I'd like tell out the environment of service, like dev, test and prod, because I have different domains for them
	log.Debug("authority: %s", authority)

	// validate service version and auth
	reqPath, err := url.Parse(path)
	if err != nil {
		return 401, errors.New("path is invalid: " + err.Error())
	}
	log.Debug("request Path: %s, QueryString: %+v", reqPath.Path, reqPath.Query())

	if reqPath.Path != chirp.Endpoint {
		return 404, errors.New("path has to be /v1/webtrans")
	}

	*userID = reqPath.Query().Get("id")
	*publicKey = reqPath.Query().Get("publickey")

	return 200, nil
}
