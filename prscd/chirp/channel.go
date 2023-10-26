// Package chirp describes peer-to-peer communication protocol.
package chirp

import (
	"sync"

	"github.com/vmihailenco/msgpack/v5"
	"github.com/yomorun/psig"
	"yomo.run/prscd/util"
)

// Channel describes a message channel.
type Channel struct {
	UniqID string   // uniq id
	AppID  string   // TODO: APP_ID
	pdic   sync.Map // all peers subscribed this channel
}

// AddPeer add peer to this channel.
func (c *Channel) AddPeer(p *Peer) {
	c.pdic.Store(p.Sid, p)
}

// RemovePeer remove peer from this channel.
func (c *Channel) RemovePeer(p *Peer) {
	c.pdic.Delete(p.Sid)
}

// Broadcast message to all peers in this channel by yomo,
// yomo create a distributed cloud network, peers from different location
// will connect to different nodes in this network, so the message will be
// broadcast to all nodes.
func (c *Channel) Broadcast(sig *psig.Signalling) {
	sigSentOverYoMo := sig.Clone()
	sigSentOverYoMo.AppID = c.AppID
	sigSentOverYoMo.MeshID = Node.MeshID
	go Node.BroadcastToYoMo(&sigSentOverYoMo)
}

// Dispatch messages to all peers in this channel of current node.
func (c *Channel) Dispatch(sig *psig.Signalling) {
	// sig.Sid is sender's sid when sending message
	log.Debug("[%s]\tSND>: %+v", sig.Sid, sig)
	var sender = sig.Sid
	// do not broadcast APP_ID and Sid to end user
	sig.AppID = ""
	sig.Sid = ""
	resp, err := msgpack.Marshal(sig)
	if err != nil {
		log.Error("msgpack marshal: %+v", err)
		return
	}

	c.pdic.Range(func(k, v interface{}) bool {
		// do not broadcast to sender-self
		sid := k.(string)
		p := v.(*Peer)
		if sid == sender {
			util.Log.Debug("-----------ignore sender-self: %s", sender)
			return true
		}
		util.Log.Debug("[%s] BroadcastPresence to ch:%s, for sid:%s", sender, c.UniqID, p.Sid)
		err = p.conn.Write(resp)
		if err != nil {
			log.Error("ws.write error: %+v", err)
		}
		return true
	})
}

// getLen returns the number of peers in this channel of current node.
func (c *Channel) getLen() int {
	var count int
	c.pdic.Range(func(k, v interface{}) bool {
		count++
		return true
	})
	return count
}
