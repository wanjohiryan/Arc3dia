package chirp

import (
	"net"

	"github.com/gobwas/ws/wsutil"
	"github.com/quic-go/quic-go"
)

// Connection is connection either WebSocket or WebTransport
type Connection interface {
	// RemoteAddr returns the client network address.
	RemoteAddr() string
	// Write the data to the connection
	Write(msg []byte) error
}

/*** WebSocket ***/

// NewWebSocketConnection creates a new WebSocketConnection
func NewWebSocketConnection(conn net.Conn) Connection {
	return &WebSocketConnection{
		underlyingConn: conn,
	}
}

// WebSocketConnection is a WebSocket connection
type WebSocketConnection struct {
	underlyingConn net.Conn
}

// RemoteAddr returns the client network address.
func (c *WebSocketConnection) RemoteAddr() string {
	return (c.underlyingConn).RemoteAddr().String()
}

// Write the data to the connection
func (c *WebSocketConnection) Write(msg []byte) error {
	return wsutil.WriteServerBinary(c.underlyingConn, msg)
}

/*** WebTransport ***/

// NewWebTransportConnection creates a new WebTransportConnection
func NewWebTransportConnection(conn quic.Connection) Connection {
	return &WebTransportConnection{
		underlyingConn: conn,
	}
}

// WebTransportConnection is a WebTransport connection
type WebTransportConnection struct {
	underlyingConn quic.Connection
}

// RemoteAddr returns the client network address.
func (c *WebTransportConnection) RemoteAddr() string {
	return c.underlyingConn.RemoteAddr().String()
}

// Write the data to the connection
func (c *WebTransportConnection) Write(msg []byte) error {
	// add 0x00 to msg
	buf := []byte{0x00}
	buf = append(buf, msg...)
	if err := c.underlyingConn.SendMessage(buf); err != nil {
		log.Error("[%s] SendMessage error: %v", c.RemoteAddr(), err)
		return err
	}
	return nil
}
