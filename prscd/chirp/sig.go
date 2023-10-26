package chirp

import "github.com/yomorun/psig"

// NewSigPeerOnline create OpPeerOnline message.
func NewSigPeerOnline(chid string, p *Peer) *psig.Signalling {
	return &psig.Signalling{
		Type:    psig.SigControl,
		OpCode:  psig.OpPeerOnline,
		Channel: chid,
		Cid:     p.Cid,
		Sid:     p.Sid,
	}
}

// NewSigPeerOffline create OpPeerOffline message.
func NewSigPeerOffline(chid string, p *Peer) *psig.Signalling {
	return &psig.Signalling{
		Type:    psig.SigControl,
		OpCode:  psig.OpPeerOffline,
		Channel: chid,
		Cid:     p.Cid,
		Sid:     p.Sid,
	}
}

// NewSigChannelJoined create OpChannelJoin message.
func NewSigChannelJoined(chName string) *psig.Signalling {
	return &psig.Signalling{
		Type:    psig.SigControl,
		OpCode:  psig.OpChannelJoin,
		Channel: chName,
		MeshID:  Node.MeshID,
	}
}
