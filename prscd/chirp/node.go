package chirp

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/vmihailenco/msgpack/v5"
	"github.com/yomorun/psig"
	"github.com/yomorun/yomo"
	"github.com/yomorun/yomo/serverless"
	"yomo.run/prscd/util"
)

var log = util.Log

const (
	// Endpoint is the base path of service
	Endpoint string = "/v1"
)

type node struct {
	cdic   sync.Map            // all channels on this node
	pdic   sync.Map            // all peers on this node
	Env    string              // Env describes the environment of this node, e.g. "dev", "prod"
	MeshID string              // MeshID describes the id of this node
	sndr   yomo.Source         // the yomo source used to send data to the geo-distributed network which built by yomo
	rcvr   yomo.StreamFunction // the yomo stream function used to receive data from the geo-distributed network which built by yomo
}

// AuthUser auth user by public key
func (n *node) AuthUser(publicKey string) (appID string, ok bool) {
	log.Info("Node| auth_user: publicKey=%s", publicKey)

	// implement your own auth logic if needed

	if n.Env == "dev" {
		log.Debug("Node| auth_user: DEV MODE, skip")
		return "DEV_APP", true
	}

	return "YOMO_APP", true
}

// AddPeer add peer to channel named `cid` on this node.
func (n *node) AddPeer(conn Connection, cid, appID string) *Peer {
	log.Info("[%s] node.add_peer: %s", conn.RemoteAddr(), cid)
	peer := &Peer{
		Sid:      conn.RemoteAddr(),
		Cid:      cid,
		Channels: make(map[string]*Channel),
		conn:     conn,
		AppID:    appID,
	}

	n.pdic.Store(n.getIDOnNode(appID, peer.Sid), peer)

	return peer
}

// RemovePeer remove peer on this node.
func (n *node) RemovePeer(appID, pid string) {
	log.Info("[%s] node.remove_peer", pid)
	n.pdic.Delete(n.getIDOnNode(appID, pid))
}

// getIDOnNode get the unique id of peer or channel on this node.
func (n *node) getIDOnNode(appID, name string) string {
	return appID + "|" + name
}

// GetOrCreateChannel get or create channel on this node.
func (n *node) GetOrAddChannel(appID, name string) *Channel {
	channelNameOnNode := n.getIDOnNode(appID, name)
	channel, ok := n.cdic.LoadOrStore(channelNameOnNode, &Channel{
		UniqID: name,
		AppID:  appID,
	})

	if !ok {
		log.Info("create channel: %s", name)
	}

	return channel.(*Channel)
}

// FindChannel returns the channel on this node by name.
func (n *node) FindChannel(appID, name string) *Channel {
	channelNameOnNode := n.getIDOnNode(appID, name)
	ch, ok := n.cdic.Load(channelNameOnNode)
	if !ok {
		log.Debug("channel not found: %s", channelNameOnNode)
		return nil
	}
	return ch.(*Channel)
}

// ConnectToYoMo connect this node to who geo-distributed network which built by yomo.
func (n *node) ConnectToYoMo(sndr yomo.Source, rcvr yomo.StreamFunction) error {
	// connect yomo source to zipper
	err := sndr.Connect()
	if err != nil {
		return err
	}

	sfnHandler := func(ctx serverless.Context) {
		var sig *psig.Signalling
		err := msgpack.Unmarshal(ctx.Data(), &sig)
		if err != nil {
			log.Error("Read from YoMo error: %v, msg=%# x, string(msg)=%s", err, ctx.Data(), ctx.Data())
		}
		log.Debug("\033[32m[\u21CA\u21CA]\t%s\033[36m", sig)

		channel := n.FindChannel(sig.AppID, sig.Channel)
		if channel != nil {
			channel.Dispatch(sig)
			log.Debug("[\u21CA]\t dispatched to %s", sig.Cid)
		} else {
			log.Debug("[\u21CA]\t dispatch to channel failed cause of not exist: %s", sig.Cid)
		}
	}

	// set observe data tags from yomo network by yomo stream function
	// 0x20 comes from other prscd nodes
	// 0x21 comes from backend sfn
	rcvr.SetObserveDataTags(0x20, 0x21)

	// handle data from yomo network, and dispatch to the same channel on this node.
	rcvr.SetHandler(sfnHandler)

	err = rcvr.Connect()
	if err != nil {
		return err
	}

	n.sndr = sndr
	n.rcvr = rcvr
	return nil
}

// BroadcastToYoMo broadcast presence to yomo
func (n *node) BroadcastToYoMo(sig *psig.Signalling) {
	// sig.Sid is sender's sid when sending message
	log.Debug("\033[34m[%s][\u21C8\u21C8]\t %s\033[36m", sig.AppID, sig)
	buf, err := msgpack.Marshal(sig)
	if err != nil {
		log.Error("msgpack marshal: %+v", err)
		return
	}

	err = n.sndr.Write(0x20, buf)
	if err != nil {
		log.Error("broadcast to yomo error: %+v", err)
	}
}

// Node describes current node, which is a singleton. There is only one node in a `prscd` process.
// But multiple `prscd` processes can be served on the same server.
var Node *node

// CreateNodeSingleton create the singleton node instance.
func CreateNodeSingleton() {
	log.Info("init Node instance, mesh_id=%s", os.Getenv("MESH_ID"))
	Node = &node{
		MeshID: os.Getenv("MESH_ID"),
	}
}

// DumpNodeState prints the user and room information to stdout.
func DumpNodeState() {
	log.Info("Dump start --------")
	Node.cdic.Range(func(k1, v1 interface{}) bool {
		log.Info("Channel:%s", k1)
		ch := v1.(*Channel)
		log.Info("\t Peers count: %d", ch.getLen())
		ch.pdic.Range(func(key, value interface{}) bool {
			log.Info("\tPeer: sid=%s, cid=%s", key, value)
			return true
		})
		return true
	})
	log.Info("Dump done --------")
}

// DumpConnectionsState prints the user and room information to stdout.
func DumpConnectionsState() {
	log.Info("Dump start --------")
	counter := make(map[string]int)
	Node.cdic.Range(func(k1, v1 interface{}) bool {
		log.Info("Channel:%s", k1)
		chName := k1.(string)
		ch := v1.(*Channel)
		peersCount := ch.getLen()
		// chName is like "appID|channelName", so we need to split it to get appID
		appID := strings.Split(chName, "|")[0]
		log.Info("\t[%s] %s Peers count: %d", appID, chName, peersCount)
		if _, ok := counter[appID]; !ok {
			counter[appID] = peersCount
		} else {
			counter[appID] += peersCount
		}
		return true
	})
	// list all counter
	for appID, count := range counter {
		log.Info("->[%s] connections: %d", appID, count)
	}
	// write counter to /tmp/conns.log
	f, err := os.OpenFile("/tmp/conns.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Error("open file: %v", err)
	}
	defer f.Close()
	timestamp := time.Now().Unix()
	for appID, count := range counter {
		if count > 0 {
			f.WriteString(fmt.Sprintf("{\"timestamp\": %d, \"conns\": %d, \"app_id\": \"%s\", \"mesh_id\": \"%s\"}\n\r", timestamp, count, appID, Node.MeshID))
		}
	}
}
