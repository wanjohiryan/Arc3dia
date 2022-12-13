<!-- inspired by https://github.com/Infisical/infisical/blob/main/README.md -->

<h3 align="center">
  <img height=120 width=300 src="./imgs/logo-black.svg#gh-light-mode-only" alt="qwantify logo">
  <img height=120 width=300 src="./imgs/logo-white.svg#gh-dark-mode-only" alt="qwantify logo">
</h3>
<p align="center">
  <p align="center">Play games with your friends right from the browser. No installations needed</p>
</p>
<h4 align="center">
  <a href="https://qwantify.vercel.app/">qwantify Arcade</a> |
  <a href="https://docs.google.com/forms/d/e/1FAIpQLSfBZDOlcdnvJwUdt5ju-v8Gx4oIqHd_jHu_p6QCvlwgaOvQ0A/viewform">Hosting</a> |
  <a href="https://qwantify.vercel.app/">Docs</a> |
  <a href="https://qwantify.vercel.app/">Website</a>
</h4>

<h4 align="center">
  <a href="https://github.com/wanjohiryan/qwantify">
    <img src="https://img.shields.io/badge/PRs-Welcome-brightgreen" alt="PRs welcome!" />
  </a>
  <a href="">
    <img src="https://img.shields.io/github/commit-activity/m/wanjohiryan/qwantify" alt="git commit activity" />
  </a>
</h4>

<pre align="center" style="width:100%;padding:50px;background: rgb(238,174,202);
background: radial-gradient(circle, rgba(238,174,202,1) 0%, rgba(148,187,233,1) 100%);">
  <img src="./imgs/web.gif" width="90%" alt="playing with qwantify"/>
</pre>

**[qwantify](https://qwantify.vercel.app)** is an open source docker image for running games (or other apps) on a shared host computer with at least one gpu.

- **User-Friendly Interface** to intuitively play games with your friends
- **Complete control over your game data** - play online, save your game progress locally
- 🛠️ **Cloud and GPU Agnostic deployment** that lets you play and host games anywhere anytime, through the browser 
- 🛠️ **Url invites for friends**
- 🛠️ **Play with multiple gamepads** per gameroom. Turn any game into multiplayer
- 🛠️ **Official Support for AMD and Intel Gpus**
- 🔜 **1-Click Deploy** locally, AWS or GCP
- 🔜 **Twitch and Youtube stream** integrations

And more.

<p align="center" >
  <a href="https://qwantify.vercel.app" target="_blank"><img src="./imgs/sign-up.png" height=120 /> </a>
  <a href="https://qwantify.vercel.app" target="_blank"><img src="./imgs/game-demo.png" height=120 /> </a>
  <a href="https://docs.google.com/forms/d/e/1FAIpQLSfBZDOlcdnvJwUdt5ju-v8Gx4oIqHd_jHu_p6QCvlwgaOvQ0A/viewform" target="_blank"><img src="./imgs/partner-up.png" height=120></a>
</p>

## 🚀 Get started

To quickly get started, pull the image and run it with docker compose (*recommended*)

```bash
version: "3.8"
services:
  qwantify:
    image: wanjohiryan/qwantify:latest #or ghcr.io/wanjohiryan/qwantify:20.04
    restart: "unless-stopped"
    ports:
      - "8080:8080" #web interface
      - "52000-52100:52000-52100/udp" #webrtc 
    volumes:
      - /games:/games #directory with folders containing your game(s)
    deploy:
      resources:
        reservations:
          devices: #share nvidia gpu (recommended)
            capabilities: [gpu] 
        limits:
          memory: 5G #depends on game (recommended is 4)
          cpus: '4' #depends on game (recommended is 4)
    environment:
      - NEKO_SCREEN=1920x1080@30 #screen size
      - NEKO_PASSWORD=neko #password for the invited guests
      - NEKO_PASSWORD_ADMIN=admin #password for the host admin 
      - NEKO_EPR=52000-52100 #webrtc ports(defaults to 52000-52100)
      - NEKO_ICELITE=1
      - APPPATH=/path/to/game/folder #folder containing the game
      - APPFILE=/game.exe #game executable file

```

Then

```bash
docker-compose up -d docker-compose.yaml
```

## 🔥 What's cool about this?

qwantify makes gaming . We're on a mission to make games more accessible to all, <i>not just gamers with expensive hardware</i>. 

We are currently working hard to make qwantify more extensive. Need any integrations or want a new feature? Feel free to [create an issue](https://github.com/wanjohiryan/qwantify/issues) or [contribute](https://github.com/wanjohiryan/qwantify/blob/master/CONTRIBUTING.md) directly to the repository.

## 🌱 Contributing

Whether it's big or small, we love contributions ❤️ Check out our guide to see how to [get started](https://github.com/wanjohiryan/qwantify/blob/master/CONTRIBUTING.md).

## 🐥 Status

- [x] Public Alpha: Anyone can sign up over at the [qwantify arcade](https://qwantify.vercel.app/) 
- [ ] Public Beta: Stable enough for most gamers.
- [ ] Public: Production-ready.

We're currently in Public Alpha.
## 🔌 Integrations

We're currently setting the foundation and building a gaming network so games can be played from anywhere on the planet. Any help is welcome! :)

<p align="center" ><b>All Available Servers<b></p>

<table align="center" width="100%" >
    <tr>
        <th scope="col">Region</th>
        <th scope="col">Instances</th>
    </tr>
    <tr>
        <th>North America</th>
        <td>
        <table>
        <tr>
            <th>Host</th>
            <th>Location</th>
            <th>Online/Offline</th>
         </tr>
         <tr>
         <th>AWS</th>
         <td>us-east-1</td>
         <td>Online ✔️</td>
         </tr>
        </table>
        </td>
    </tr>
    <tr>
        <th scope="row">Africa</th>
        <td><table>
        <tr>
            <th>Host</th>
            <th>Location</th>
            <th>Online/Offline</th>
         </tr>
         <tr>
         <th>AWS</th>
         <td>af-south-2</td>
         <td>Online ✔️</td>
         </tr>
        </table></td>
    </tr>
     <tr>
        <th scope="row">Europe</th>
       <td><table>
        <tr>
            <th>Host</th>
            <th>Location</th>
            <th>Online/Offline</th>
         </tr>
         <tr>
         <th>Indie</th>
         <td>Berlin</td>
         <td>Online ✔️</td>
         </tr>
        </table></td>
    </tr>
    <tr>
        <th scope="row">South America</th>
        <td>Currently not available :( ❌</td>
    </tr>
    <tr>
        <th scope="row">Asia</th>
        <td>Currently not available :( ❌</td>
    </tr>
</table>

### Rent out your GPU and help us deliver games to everyone.

<p >
  <a href="https://docs.google.com/forms/d/e/1FAIpQLSfBZDOlcdnvJwUdt5ju-v8Gx4oIqHd_jHu_p6QCvlwgaOvQ0A/viewform" target="_blank"><img src="./imgs/partner-up.png" height=120></a>
</p>



**Stay frosty :)**