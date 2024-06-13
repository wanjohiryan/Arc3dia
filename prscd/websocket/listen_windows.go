//go:build windows
// +build windows

package websocket

import (
	"net"
)

var lc = net.ListenConfig{}
