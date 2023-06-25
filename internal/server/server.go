package server

import (
	"context"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/wanjohiryan/Arc3dia/internal/warp"

	"github.com/kixelated/invoker"
	"github.com/kixelated/quic-go"
	"github.com/kixelated/quic-go/http3"
	"github.com/kixelated/quic-go/logging"
	"github.com/kixelated/quic-go/qlog"
	"github.com/kixelated/webtransport-go"
)

// To track whether game is running or not
var isRunning bool
var mutex sync.Mutex

type Server struct {
	inner    *webtransport.Server
	media    *warp.Media
	sessions invoker.Tasks
	cert     *tls.Certificate
	// gamePath string
}

type Config struct {
	Addr   string
	Cert   *tls.Certificate
	LogDir string
	Media  *warp.Media
	// Game   string
}

func New(config Config) (s *Server, err error) {
	s = new(Server)
	s.cert = config.Cert
	s.media = config.Media
	// s.gamePath = config.Game

	quicConfig := &quic.Config{}

	if config.LogDir != "" {
		quicConfig.Tracer = qlog.NewTracer(func(p logging.Perspective, connectionID []byte) io.WriteCloser {
			path := fmt.Sprintf("%s-%s.qlog", p, hex.EncodeToString(connectionID))

			f, err := os.Create(filepath.Join(config.LogDir, path))
			if err != nil {
				// lame
				panic(err)
			}

			return f
		})
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{*s.cert},
	}

	// Host a HTTP/3 server to serve the WebTransport endpoint
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.handleWatch)
	//TODO: come back to this and fix it
	// mux.HandleFunc("/play", s.handlePlay)
	//TODO: add state share, state snapshot here
	// mux.HandleFunc("/game", s.handleGame)

	s.inner = &webtransport.Server{
		H3: http3.Server{
			TLSConfig:  tlsConfig,
			QuicConfig: quicConfig,
			Addr:       config.Addr,
			Handler:    mux,
		},
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	return s, nil
}

func (s *Server) runServe(ctx context.Context) (err error) {
	return s.inner.ListenAndServe()
}

func (s *Server) runShutdown(ctx context.Context) (err error) {
	<-ctx.Done()
	s.inner.Close() // close on context shutdown
	return ctx.Err()
}

func (s *Server) Run(ctx context.Context) (err error) {
	return invoker.Run(ctx, s.runServe, s.runShutdown, s.sessions.Repeat)
}

func (s *Server) handleWatch(w http.ResponseWriter, r *http.Request) {
	hijacker, ok := w.(http3.Hijacker)
	if !ok {
		panic("unable to hijack connection: must use kixelated/quic-go")
	}

	conn := hijacker.Connection()

	sess, err := s.inner.Upgrade(w, r)
	if err != nil {
		http.Error(w, "failed to upgrade session", 500)
		return
	}

	err = s.serveWatchSession(r.Context(), conn, sess)
	if err != nil {
		log.Println(err)
	}
}

// TODO: upgrade to webtransport?
// func (s *Server) handleGame(w http.ResponseWriter, r *http.Request) {
// 	mutex.Lock()
// 	defer mutex.Unlock()

// 	if s.gamePath == "" {
// 		fmt.Println("No entrypoint script provided")
// 		http.Error(w, "No entrypoint script provided", 500)
// 		return
// 	}

// 	if !isRunning {
// 		// Check if entrypoint for the game exists
// 		// entrypoint, exists := os.LookupEnv("ENTRYPOINT")
// 		// if !exists {
// 		// 	fmt.Println("ENTRYPOINT environment variable not set")
// 		// 	http.Error(w, "ENTRYPOINT env variable not found", 500)
// 		// 	return
// 		// }

// 		var entrypoint string

// 		if s.gamePath != "" {
// 			entrypoint = s.gamePath
// 		} else {
// 			gamePath, exists := os.LookupEnv("ENTRYPOINT")
// 			if !exists {
// 				fmt.Println("ENTRYPOINT environment variable not set")

// 				w.WriteHeader(http.StatusInternalServerError)
// 				w.Write([]byte("ENTRYPOINT environment variable not set"))
// 				return
// 			}
// 			entrypoint = gamePath
// 		}

// 		if _, err := os.Stat(entrypoint); err == nil {
// 			fmt.Println("Bash script not found")
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte("ENTRYPOINT bash script not found"))
// 			return
// 		}

// 		cmd := exec.Command("bash", entrypoint)

// 		output, err := cmd.CombinedOutput()
// 		if err != nil {
// 			fmt.Println("Error running game:", err)
// 			http.Error(w, "Error running game", 500)
// 			return
// 		}

// 		w.Write([]byte("Game running successfully"))
// 		w.WriteHeader(http.StatusOK)

// 		fmt.Println("Game started successfully:", string(output))

// 		isRunning = true
// 		return
// 	} else {
// 		w.Write([]byte("Game already running"))
// 		w.WriteHeader(http.StatusOK)
// 		return
// 	}
// }

func (s *Server) serveWatchSession(ctx context.Context, conn quic.Connection, sess *webtransport.Session) (err error) {
	defer func() {
		if err != nil {
			sess.CloseWithError(1, err.Error())
		} else {
			sess.CloseWithError(0, "end of broadcast")
		}
	}()

	ss, err := warp.NewSession(conn, sess, s.media)
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}

	err = ss.Run(ctx)
	if err != nil {
		return fmt.Errorf("terminated session: %w", err)
	}

	return nil
}

// func (s *Server) handlePlay(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("new Player has connected")

// 	hijacker, ok := w.(http3.Hijacker)
// 	if !ok {
// 		panic("unable to hijack connection: must use kixelated/quic-go")
// 	}

// 	conn := hijacker.Connection()

// 	sess, err := s.inner.Upgrade(w, r)
// 	if err != nil {
// 		http.Error(w, "failed to upgrade session", 500)
// 		return
// 	}

// 	err = s.servePlaySession(r.Context(), conn, sess)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func (s *Server) servePlaySession(ctx context.Context, conn quic.Connection, sess *webtransport.Session) (err error) {
// 	defer func() {
// 		if err != nil {
// 			sess.CloseWithError(1, err.Error())
// 		} else {
// 			sess.CloseWithError(0, "end of broadcast")
// 		}
// 	}()

// 	ss, err := play.NewSession(conn, sess)
// 	if err != nil {
// 		return fmt.Errorf("failed to create session: %w", err)
// 	}

// 	err = ss.Run(ctx)
// 	if err != nil {
// 		return fmt.Errorf("terminated session: %w", err)
// 	}

// 	return nil
// }
