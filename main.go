package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"

	"github.com/kixelated/invoker"
	"github.com/wanjohiryan/Arc3dia/internal/server"
)

//FIXME: add media, game and certs

func main() {
	err := run(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) (err error) {
	addr := flag.String("port", ":8080", "HTTPS server address")
	cert := flag.String("cert", "./cert/localhost.crt", "TLS certificate file path")
	key := flag.String("key", "./cert/localhost.key", "TLS certificate file path")
	logDir := flag.String("log-dir", "", "logs will be written to the provided directory")

	// dash := flag.String("dash", "./media/playlist.mpd", "DASH playlist path")
	// game := flag.String("game", "", "path to game executable")

	flag.Parse()

	// media, err := warp.NewMedia(*dash)
	// if err != nil {
	// 	return fmt.Errorf("failed to open media: %w", err)
	// }

	tlsCert, err := tls.LoadX509KeyPair(*cert, *key)
	if err != nil {
		return fmt.Errorf("failed to load TLS certificate: %w", err)
	}

	warpConfig := server.Config{
		Addr:   *addr,
		Cert:   &tlsCert,
		LogDir: *logDir,
		// Media:  media,
		// Game:   *game,
	}

	fmt.Print(warpConfig)

	// warpServer, err := server.New(warpConfig)
	// if err != nil {
	// 	return fmt.Errorf("failed to create warp server: %w", err)
	// }

	log.Printf("listening on %s", *addr)
	//warpServer.Run

	return invoker.Run(ctx, invoker.Interrupt)
}
