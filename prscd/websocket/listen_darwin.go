//go:build darwin
// +build darwin

package websocket

import (
	"net"
	"syscall"

	"golang.org/x/sys/unix"
)

// reuse port on darwin
var lc = net.ListenConfig{
	Control: func(network, address string, c syscall.RawConn) error {
		var opErr error
		if err := c.Control(func(fd uintptr) {
			// 端口复用，这样可以多进程监听该端口，充分利用 CPU 资源；同时也可以实现热更新
			opErr = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
		}); err != nil {
			return err
		}
		return opErr
	},
}
