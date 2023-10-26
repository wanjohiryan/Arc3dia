// package main start the service
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"time"

	"yomo.run/prscd/chirp"
	"yomo.run/prscd/util"
	"yomo.run/prscd/websocket"
	"yomo.run/prscd/webtransport"

	"github.com/joho/godotenv"
	"github.com/yomorun/yomo"
	"github.com/yomorun/yomo/pkg/config"
	"github.com/yomorun/yomo/pkg/trace"
)

var log = util.Log

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	chirp.CreateNodeSingleton()

	// DEBUG env indicates development mode, verbose log
	if os.Getenv("DEBUG") == "true" {
		chirp.Node.Env = "development"
		log.SetLogLevel(util.DEBUG)
		log.Debug("IN DEVELOPMENT ENV")
	}

	// AS_YOMO_ZIPPER env indicates start YOMO Zipper in this process
	if os.Getenv("AS_YOMO_ZIPPER") == "true" {
		go startYomoZipper()
		// sleep 2 seconds to wait for YoMo Zipper ready
		time.Sleep(2 * time.Second)
	} else {
		log.Debug("Skip start YOMO Zipper")
	}

	// YOMO_ZIPPER env indicates the endpoint of YoMo Zipper to connect
	log.Debug("connect to YoMo Zipper: %s", os.Getenv("YOMO_ZIPPER"))

	// add open tracing
	tp, shutdown, err := trace.NewTracerProviderWithJaeger("prscd")
	if err == nil {
		log.Info("[%s] 🛰 trace enabled", "prscd")
	}
	defer shutdown(context.Background())

	// sndr is sender to send data to other prscd nodes by YoMo
	sndr := yomo.NewSource(
		os.Getenv("YOMO_SNDR_NAME"),
		os.Getenv("YOMO_ZIPPER"),
		yomo.WithCredential(os.Getenv("YOMO_CREDENTIAL")),
		yomo.WithTracerProvider(tp),
	)

	// rcvr is receiver to receive data from other prscd nodes by YoMo
	rcvr := yomo.NewStreamFunction(
		os.Getenv("YOMO_RCVR_NAME"),
		os.Getenv("YOMO_ZIPPER"),
		yomo.WithSfnCredential(os.Getenv("YOMO_CREDENTIAL")),
		yomo.WithSfnTracerProvider(tp),
	)

	// connect to YoMo
	chirp.Node.ConnectToYoMo(sndr, rcvr)

	// default addr and port listening
	addr := "0.0.0.0:443"
	if os.Getenv("PORT") != "" {
		addr = fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))
	}

	// load TLS cert and key, halt if error occurs,
	// this helped developers to find out TLS related issues asap.
	config, err := loadTLS(os.Getenv("CERT_FILE"), os.Getenv("KEY_FILE"), os.Getenv("DOMAIN"))
	if err != nil {
		log.Fatal(err)
		os.Exit(-2)
	}

	// start WebSocket listener
	go websocket.ListenAndServe(addr, config)

	// start WebTransport listener
	go webtransport.ListenAndServe(addr, config)

	// start Probe Server for AWS health check
	go startProbeServer(61226)

	// Ctrl-C or kill <pid> graceful shutdown
	// - `kill -SIGUSR1 <pid>` customize
	// - `kill -SIGTERM <pid>` graceful shutdown
	// - `kill -SIGUSR2 <pid>` inspect golang GC
	log.Info("PID: %d", os.Getpid())
	// write pid to ./prscd.pid, overwrite if exists
	pidFile := "./prscd.pid"
	f, err := os.OpenFile(pidFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%d", os.Getpid()))
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("Prscd Dev Server is running on https://%s:%s/v1", os.Getenv("DOMAIN"), os.Getenv("PORT"))

	c := make(chan os.Signal, 1)
	registerSignal(c)
}

func startYomoZipper() {
	conf, err := config.ParseConfigFile("./yomo.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("integrated YoMo config: %v", conf)
	log.Debug("integrated YoMo zipper: %s", fmt.Sprintf("%s:%d", conf.Host, conf.Port))

	zipper, err := yomo.NewZipper(conf.Name, conf.Functions, conf.Downstreams)
	if err != nil {
		log.Fatal(err)
	}

	err = zipper.ListenAndServe(context.Background(), fmt.Sprintf("%s:%d", conf.Host, conf.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func loadTLS(certFile, keyFile, domain string) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	// check if TLS cert is expired
	// Parse the X.509 certificate
	parsedCert, err := x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return nil, err
	}

	// Get the expiration date
	expirationDate := parsedCert.NotAfter
	log.Debug("check TLS cert expiration date: %s", expirationDate)

	// determine if the certificate is expired
	if time.Now().After(expirationDate) {
		return nil, fmt.Errorf("tls cert is expired")
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		NextProtos:   []string{"http/1.1", "h2", "h3", "http/0.9", "http/1.0", "spdy/1", "spdy/2", "spdy/3"},
	}, nil
}
