
# 🌎 Prscd

The open source backend for Presencejs v2.0

## 🎯 Roadmap

- [x] Websocket arraybuffer support
- [x] Zero-copy upgrade to WebSocket
- [x] SO_REUSEPORT on Darwin and Linux
- [x] Implement WebSocket native Ping/Pong frame to keep alive
- [ ] reuse goroutine
- [x] WebTransport Datagram support, unreliable but fast communication
- [ ] WebTransport Stream support, reliable
- [ ] pprof
- [x] Prscd Clusters by YoMo
- [x] Geo-distributed System / Distributed Cloud arch by YoMo

## 🥷🏻 Development

1. Start prscd service in terminal-2：`make dev`
1. Open `webtransport.html` by Chrome with Dev Tools
1. Open `websocket.html` by Chrome with Dev Tools

![](https://github.com/fanweixiao/gifs-repo/blob/main/prscd-readme.gif)

[![asciicast](https://asciinema.org/a/565542.svg)](https://asciinema.org/a/565542)

## 🦸🏻 Self-hosting

Compile:

```bash
make dist
```

### ☝🏻 Host on Single Cloud Region

TODO: how to deploy `prscd` on digitalocean

TODO: introducing [geoping.gg](https://geoping.gg), lighthouse for realtime applications.

### 🌍 Host as Geo-distributed System

DNS improvements explaination

#### by Vercel edge functions

Redirect end-user connect to node close to them.

#### by AWS Global Accelarator

Anycast IP

#### by Azure Traffic Manager

Geo-IP

## ☕️ FAQ

### about https://lo.yomo.dev

```bash
$ openssl x509 -enddate -noout -in prscd/lo.yomo.dev.cert
notAfter=May 22 07:40:45 2023 GMT
```

### how to generate SSL for your own domain

1. `brew install certbot`
2. `sudo certbot certonly --manual --preferred-challenges dns -d prscd.example.com`
3. create a TXT record followed the instruction by certbot
4. `nslookup -type=TXT _acme-challenge.prscd.example.com` to verify the process
5. `sudo chown -Rv "$(whoami)":staff /etc/letsencrypt/` set permission
6. cert and key: `/etc/letsencrypt/live/prscd.example.com/{fullchain, privkey}.pem`
7. verify the expiratioin time: `openssl x509 -enddate -noout -in prscd.example.com.cert.pem`

### if you are behind a proxy on Mac

Most of proxy applications drop WebTransport or HTTP/3, so if you are a macos user,
this bash script can helped bypass `*.yomo.dev` domain to proxy.

```bash
networksetup -setproxybypassdomains "Wi-Fi" $(networksetup -getproxybypassdomains "Wi-Fi" | awk '{ printf "\"%s\" ", $0 }') "*.yomo.dev"
```

### Integrate to your own Auth system

Currently, provide `public_key` for authentication, the endpoint looks like: `/v1?app_id=<USER_CLIENT_ID>&public_key=<PUBLIC_KEY>`

### Live inspection

Execute `make dev` in terminal-1:

```bash
$ make dev
go run -race main.go
pid: 20079
Listening SIGUSR1, SIGUSR2, SIGTERM/SIGINT...
```

Open terminal-2, execute:

```bash
$ kill -SIGUSR1 20079
$ kill -SIGUSR2 20079
```

The output of terminal-1 will looks like:

```bash
$ make dev
go run -race main.go
pid: 20079
Listening SIGUSR1, SIGUSR2, SIGTERM/SIGINT...
Received signal: user defined signal 1
SIGUSR1
Dump start --------
Peers: 1
Channel:room-1
	Peer:127.0.0.1:62577
Dump doen --------
Received signal: user defined signal 2
	NumGC = 0
```

### Configure Firewall of Cloud Provider

TCP and UDP on the `PORT` shall has to be allowed in security rules.

## .env File

- `DEBUG=true`: debug mode
- `PORT=443`: indicate the PORT used to listen, both WebSocket and WebTransport
- `MESH_ID=MID_EAST`: indicate nodes in distributed cloud archtecture
- `YOMO_SNDR_NAME`: the name of YoMo Source
- `YOMO_RCVR_NAME`: the name of YoMo Stream Function
- `CERT_FILE`: The SSL cert file path of prscd
- `YOMO_TRACE_JAEGER_ENDPOINT`: Jaeger collector endpoint, e.g., http://localhost:14268/api/traces

- `KEY_FILE`: The SSL key file path of prscd
