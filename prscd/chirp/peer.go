package chirp

import (
	"errors"
	"io"
	"sync"

	"github.com/vmihailenco/msgpack/v5"
	"github.com/yomorun/psig"
)

// Peer describes user on this node.
type Peer struct {
	// Sid describes the unique id of this peer on this node, only used for backend.
	Sid string
	// Cid describes the unique id of this peer on who geo-distributed network, set by developer.
	Cid string
	// Channel describes the channel which this peer joined.
	Channels map[string]*Channel
	// conn is the connection of this peer.
	conn Connection
	mu   sync.Mutex
	// AppID is the id of the app which this peer belongs to.
	AppID string
}

// Join this peer to channel named `channelName`.
func (p *Peer) Join(channelName string) {
	// find channel on this node, if not exist, create it.
	c := Node.GetOrAddChannel(p.AppID, channelName)

	// add peer to this channel
	c.AddPeer(p)

	// and this channel to peer's channel list
	p.Channels[channelName] = c

	// ACK to peer has joined
	p.NotifyBack(NewSigChannelJoined(channelName))

	log.Info("[%s] ack peer.join_chanel:%s, cid=%s", p.Sid, c.UniqID, p.Cid)
}

// NotifyBack to peer with message.
func (p *Peer) NotifyBack(sig *psig.Signalling) {
	resp, err := msgpack.Marshal(sig)
	if err != nil {
		log.Error("msgpack marshal: %+v", err)
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	err = p.conn.Write(resp)

	if err != nil {
		log.Error("NotifyBack error: %+v", err)
	}
	log.Debug("[%s]\tSND>: %s", p.Sid, sig)
}

// Leave a channel
func (p *Peer) Leave(channelName string) {
	// remove channel from peer's channel list
	p.mu.Lock()
	delete(p.Channels, channelName)
	p.mu.Unlock()

	// remove peer from channel's peer list
	c := Node.FindChannel(p.AppID, channelName)
	if c == nil {
		log.Error("peer.Leave(): channel is nil. pid: %s, channel: %s", p.Sid, channelName)
		return
	}

	c.RemovePeer(p)

	// Notify others on this channel that this peer has left
	c.Broadcast(NewSigPeerOffline(channelName, p))
	log.Info("[%s] peer.leave: %s", p.Sid, c.UniqID)
}

// Disconnect clears resources of this peer when leave.
func (p *Peer) Disconnect() {
	log.Info("[%s] peer.disconnect", p.Sid)
	// wipe this peer from all channels joined before
	for _, ch := range p.Channels {
		p.Leave(ch.UniqID)
	}
	// wipe this peer from current node
	Node.RemovePeer(p.AppID, p.Sid)
}

// BroadcastToChannel will broadcast message to channel.
func (p *Peer) BroadcastToChannel(sig *psig.Signalling) {
	sig.Cid = p.Cid
	c := p.Channels[sig.Channel]
	if c == nil {
		log.Error("BroadcastToChannel: channel=%s is nil, should panic here", sig.Channel)
		return
	}

	c.Broadcast(sig)
}

// HandleSignal handle message sent from connection.
func (p *Peer) HandleSignal(r io.Reader) error {
	decoder := msgpack.NewDecoder(r)
	sig := &psig.Signalling{}
	if err := decoder.Decode(sig); err != nil {
		log.Error("msgpack.decode err, ignore: %+v", err)
		return err
	}

	// p.Sid is the id of connection, set by backend.
	sig.Sid = p.Sid
	log.Debug("[%s] >RCV: %v", p.Sid, sig)

	if sig.Type == psig.SigControl {
		// handle the Control Signalling
		switch sig.OpCode {
		case psig.OpChannelJoin: // `channel_join` signalling
			// join channel
			p.Join(sig.Channel)
		case psig.OpState: // `peer_state` signalling
			// Alice can notify Bob that her state has been updated, also,
			// Bob can use this signalling to initialize or update Alice's state
			if sig.Sid != "" && sig.Cid != "" {
				// if peer sid and client id are both set, then update the client id of this peer
				p.Cid = sig.Cid
				log.Info("Peer: %s state new ClientID: %s", p.Sid, p.Cid)
			}
			p.BroadcastToChannel(sig)
		case psig.OpPeerOffline: // `peer_offline` signalling
			p.Leave(sig.Channel)
		case psig.OpPeerOnline: // `peer_online` signalling
			p.BroadcastToChannel(sig)
		default:
			log.Error("Unknown control opcode: %d", sig.OpCode)
		}
	} else if sig.Type == psig.SigData {
		// handle the Data Signalling
		p.BroadcastToChannel(sig)
	} else {
		log.Error("ILLEGAL sig.Type, should be `data` or `control`: %+v", sig)
		return errors.New("ILLEGAL sig.Type, should be `data` or `control`")
	}

	return nil
}
