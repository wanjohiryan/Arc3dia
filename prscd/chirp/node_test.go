package chirp

import (
	"testing"

	"github.com/yomorun/yomo/core/frame"
	"yomo.run/prscd/util"
)

// NewMockConnection creates a new WebSocketConnection
func NewMockConnection(sid string) Connection {
	return &MockConnection{
		sid: sid,
	}
}

// MockConnection is a WebSocket connection
type MockConnection struct {
	sid string
}

// RemoteAddr returns the client network address.
func (c *MockConnection) RemoteAddr() string {
	return c.sid
}

// Write the data to the connection
func (c *MockConnection) Write(msg []byte) error {
	return nil
}

// SenderMock implement yomo.Source interface
type SenderMock struct{}

func (s *SenderMock) Close() error {
	return nil
}

func (s *SenderMock) Connect() error {
	return nil
}

func (s *SenderMock) Write(tag frame.Tag, data []byte) error {
	return nil
}

func (s *SenderMock) SetErrorHandler(fn func(err error)) {
}

func (s *SenderMock) SetReceiveHandler(fn func(tag frame.Tag, data []byte)) {
}

func (s *SenderMock) Broadcast(tag uint32, data []byte) error {
	return nil
}

func (s *SenderMock) SetDataTag(tag frame.Tag) {}

var channelName, appID, peerName string

func init() {
	CreateNodeSingleton()

	// mock YoMo Source
	Node.sndr = &SenderMock{}

	channelName = "test_channel"
	appID = "test_appid"
	peerName = "test_peer"

	// error level
	util.Log.SetLogLevel(2)
}

func Test_node_AddPeer(t *testing.T) {
	peer := Node.AddPeer(NewMockConnection(peerName), channelName, appID)
	peer.Join(channelName)

	assert(t, peer != nil, "peer should not be nil")
	assert(t, peer.AppID == appID, "peer.AppID should be %s, but got %s", appID, peer.AppID)
	assert(t, peer.Channels != nil, "peer.Channels should not be nil")
	assert(t, len(peer.Channels) == 1, "len(peer.Channels) should be 1, but got %d", len(peer.Channels))
	assert(t, peer.Channels[channelName] != nil, "peer.Channels[%s] should not be nil", channelName)
	ch := Node.FindChannel(appID, channelName)
	assert(t, ch != nil, "node.cdic[%s] should not be nil", appID+"|"+channelName)
	assert(t, ch.getLen() > 0, "len(node.cdic[%s].peers) should > 0", appID+"|"+channelName)
	p, ok := Node.pdic.Load(appID + "|" + peerName)
	assert(t, ok, "node.pdic[%s] should not be nil", appID+"|"+peerName)
	assert(t, p.(*Peer).Sid == peerName, "node.pdic[%s] should not be nil", appID+"|"+peerName)

	peer.Leave(channelName)
	assert(t, len(peer.Channels) == 0, "len(peer.Channels) should be 1, but got %d", len(peer.Channels))
	ch = Node.FindChannel(appID, channelName)
	assert(t, ch != nil, "node.cdic[%s] should not be nil", appID+"|"+channelName)
	assert(t, ch.getLen() == 0, "len(node.cdic[%s].pdic) should be 0, but got %d", appID+"|"+channelName, ch.getLen())
	p, ok = Node.pdic.Load(appID + "|" + peerName)
	assert(t, ok, "node.pdic[%s] should not be nil", appID+"|"+peerName)
	assert(t, p.(*Peer).Sid == peerName, "node.pdic[%s] should not be nil", appID+"|"+peerName)

	peer.Disconnect()
	ch = Node.FindChannel(appID, channelName)
	assert(t, ch != nil, "node.cdic[%s] should not be nil", appID+"|"+channelName)
	assert(t, ch.getLen() == 0, "len(node.cdic[%s].pdic) should be 0, but got %d", appID+"|"+channelName, ch.getLen())
	p, ok = Node.pdic.Load(appID + "|" + peerName)
	assert(t, !ok, "node.pdic[%s] should be nil", appID+"|"+peerName)
	assert(t, p == nil, "node.pdic[%s] should not be nil", appID+"|"+peerName)
}

func assert(t *testing.T, condition bool, format string, args ...any) {
	if !condition {
		t.Errorf(format, args...)
	}
}

func BenchmarkPeerJoinAndLeave(b *testing.B) {
	for i := 0; i < b.N; i++ {
		peer := Node.AddPeer(NewMockConnection(peerName), channelName, appID)
		peer.Join(channelName)
		peer.Leave(channelName)
		peer.Disconnect()
	}
}

func Test_node_AuthUser(t *testing.T) {
	var wantAppID = "YOMO_APP"
	gotAppID, gotOk := Node.AuthUser("kmJAUnCtkWbkNnhXYtZAGEJzGDGpFo1e1vkp6cm")
	if gotAppID != wantAppID {
		t.Errorf("node.AuthUser() gotAppID = %v, want %v", gotAppID, wantAppID)
	}
	if gotOk != true {
		t.Errorf("node.AuthUser() gotOk = %v, want %v", gotOk, true)
	}
}
