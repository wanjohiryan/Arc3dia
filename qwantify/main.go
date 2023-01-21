package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"

	"github.com/kixelated/invoker"
	"github.com/wanjohiryan/time-warp/internal/warp"
)

func main() {
	err := run(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) (err error) {
	addr := flag.String("addr", ":8080", "HTTPS server address")
	cert := flag.String("tls-cert", "./certs/localhost.crt", "TLS certificate file path")
	key := flag.String("tls-key", "./certs/localhost.key", "TLS certificate file path")
	logDir := flag.String("log-dir", "", "logs will be written to the provided directory")

	dash := flag.String("dash", "./media/playlist.mpd", "DASH playlist path")

	flag.Parse()

	media, err := warp.NewMedia(*dash)
	if err != nil {
		return fmt.Errorf("failed to open media: %w", err)
	}

	tlsCert, err := tls.LoadX509KeyPair(*cert, *key)
	if err != nil {
		return fmt.Errorf("failed to load TLS certificate: %w", err)
	}

	config := warp.ServerConfig{
		Addr:   *addr,
		Cert:   &tlsCert,
		LogDir: *logDir,
	}

	ws, err := warp.NewServer(config, media)
	if err != nil {
		return fmt.Errorf("failed to create warp server: %w", err)
	}

	log.Printf("listening on %s", *addr)

	return invoker.Run(ctx, invoker.Interrupt, ws.Run)
}
