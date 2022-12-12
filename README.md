<h1 align="center">
  <img width=300 src="/imgs/logo-black.svg#gh-light-mode-only" alt="qwantify logo">
  <img width=300 src="/imgs/logo-white.svg#gh-dark-mode-only" alt="qwantify logo">
</h1>
<p align="center">
  <p align="center">Play games with your friends right from the browser. No installations needed</p>
</p>
<h4 align="center">
  <a href="https://qwantify.vercel.app/">qwantify Arcade</a> |
  <a href="https://qwantify.vercel.app/">Self Hosting</a> |
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

<img src="/img/infisical_github_repo.png" width="100%" alt="Dashboard"/>

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
  <a href="https://qwantify.vercel.app" target="_blank"><img src="/imgs/sign-up.png" height=120 /> </a>
  <a href="https://qwantify.vercel.app" target="_blank"><img src="/imgs/demo-game.png" height=120 /> </a>
  <a href="" target="_blank"><img src="/imgs/partner-up.png" height=120></a>
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

We are currently working hard to make qwantify more extensive. Need any integrations or want a new feature? Feel free to [create an issue](https://github.com/wanjohiryan/qwantify/issues) or [contribute](https://infisical.com/docs/contributing/overview) directly to the repository.

## 🌱 Contributing

Whether it's big or small, we love contributions ❤️ Check out our guide to see how to [get started](https://infisical.com/docs/contributing/overview).

## 🐥 Status

- [x] Public Alpha: Anyone can sign up over at the [qwantify arcade](https://qwantify.vercel.app/) 
- [ ] Public Beta: Stable enough for most gamers.
- [ ] Public: Production-ready.

We're currently in Public Alpha.
## 🔌 Integrations

We're currently setting the foundation and building a gaming network so games can be played from anywhere on the planet. Any help is welcome! :)

<table align="center" width="100%" >
    <caption align="center" ><b>All Available Servers<b></caption>
    <tr>
        <th scope="col">Region</th>
        <th scope="col">Instances</th>
    </tr>
    <tr>
        <th>North America</th>
        <td>
        <table>
        <tr>
            <th>Provider</th>
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
            <th>Provider</th>
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
            <th>Provider</th>
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

### Rent out your GPU and help us deliver games to everyone. Become a partner today.

<p align="center" >
  <a href="" target="_blank"><img src="/imgs/partner-up.png" height=120></a>
</p>

## 🏘 Open-source vs. paid

This repo is entirely MIT licensed, with the exception of the `ee` directory which will contain premium enterprise features requiring a Infisical license in the future. We're currently focused on developing non-enterprise offerings first that should suit most use-cases.

## 🛡 Security

Looking to report a security vulnerability? Please don't post about it in GitHub issue. Instead, refer to our [SECURITY.md](./SECURITY.md) file.

## 🦸 Contributors

[//]: contributor-faces

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->

<a href="https://github.com/dangtony98"><img src="https://avatars.githubusercontent.com/u/25857006?v=4" width="50" height="50" alt=""/></a> <a href="https://github.com/mv-turtle"><img src="https://avatars.githubusercontent.com/u/78047717?s=96&v=4" width="50" height="50" alt=""/></a> <a href="https://github.com/maidul98"><img src="https://avatars.githubusercontent.com/u/9300960?v=4" width="50" height="50" alt=""/></a> <a href="https://github.com/gangjun06"><img src="https://avatars.githubusercontent.com/u/50910815?v=4" width="50" height="50" alt=""/></a> <a href="https://github.com/reginaldbondoc"><img src="https://avatars.githubusercontent.com/u/7693108?v=4" width="50" height="50" alt=""/></a> <a href="https://github.com/SH5H"><img src="https://avatars.githubusercontent.com/u/25437192?v=4" width="50" height="50" alt=""/></a> <a href="https://github.com/asharonbaltazar"><img src="https://avatars.githubusercontent.com/u/58940073?v=4" width="50" height="50" alt=""/></a> <a href="https://github.com/edgarrmondragon"><img src="https://avatars.githubusercontent.com/u/16805946?v=4" width="50" height="50" alt=""/></a> <a href="https://github.com/hanywang2"><img src="https://avatars.githubusercontent.com/u/44352119?v=4" width="50" height="50" alt=""/></a> <a href="https://github.com/tobias-mintlify"><img src="https://avatars.githubusercontent.com/u/110702161?v=4" width="50" height="50" alt=""/></a> <a href="https://github.com/0xflotus"><img src="https://avatars.githubusercontent.com/u/26602940?v=4" width="50" height="50" alt=""/></a> 
