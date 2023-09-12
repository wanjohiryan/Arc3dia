package main

// implement a tcp server that accept netcat request, use as a probe server for AWS health check
import (
	"fmt"
	"io"
	"net"
)

// startProbeServer create a tcp listener on port, accept command:
// nc -z -v -w1 lo.yomo.dev 61226 2>&1 |grep succeeded
func startProbeServer(port int) {
	// launch TCP server
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))

	if err != nil {
		log.Error("can not start probe server: %v", err)
		return
	}

	log.Info("Listening for connections on %s", listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error("Error accepting connection from client: %s", err)
		} else {
			go (func(conn net.Conn) {
				_, err := io.ReadAll(conn)
				if err != nil {
					log.Error("probe server process error: %v", err)
					conn.Close()
				} else {
					log.Inspect("probed from %s", conn.RemoteAddr().String())
				}
			})(conn)
		}
	}
}
