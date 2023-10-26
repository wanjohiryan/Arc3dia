//go:build !windows
// +build !windows

package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"yomo.run/prscd/chirp"
)

func registerSignal(c chan os.Signal) {
	signal.Notify(c, syscall.SIGTERM, syscall.SIGUSR2, syscall.SIGUSR1, syscall.SIGINT)
	log.Info("Listening SIGUSR1, SIGUSR2, SIGTERM/SIGINT...")
	for p1 := range c {
		log.Info("Received signal: %s", p1)
		if p1 == syscall.SIGTERM || p1 == syscall.SIGINT {
			log.Info("graceful shutting down ... %s", p1)
			os.Exit(0)
		} else if p1 == syscall.SIGUSR2 {
			// kill -SIGUSR2 <pid> will write ystat logs to /tmp/conns.log
			chirp.DumpConnectionsState()
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			fmt.Printf("\tNumGC = %v\n", m.NumGC)
		} else if p1 == syscall.SIGUSR1 {
			log.Info("SIGUSR1")
			chirp.DumpNodeState()
		}
	}
}
