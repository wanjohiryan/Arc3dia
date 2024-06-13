//go:build freebsd || linux || netbsd || openbsd
// +build freebsd linux netbsd openbsd

package websocket

import (
	"golang.org/x/sys/unix"
	"net"
	"syscall"
)

var lc = net.ListenConfig{
	Control: func(network, address string, c syscall.RawConn) error {
		var opErr error
		if err := c.Control(func(fd uintptr) {
			// reuse port
			opErr = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
			// TCP_NODELAY
			opErr = unix.SetsockoptInt(int(fd), unix.IPPROTO_TCP, unix.TCP_NODELAY, 1)
			// set priority of the socket
			opErr = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_PRIORITY, 6)
		}); err != nil {
			return err
		}
		return opErr
	},
}
