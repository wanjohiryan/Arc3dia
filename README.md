<!-- inspired by https://github.com/Infisical/infisical/blob/main/README.md -->

<!-- !TODO: Requires better description of what we do, cut on the sales lingo-->

<h3 align="center">
  <img height=120 width=300 src="./imgs/logo-black.svg#gh-light-mode-only" alt="qwantify logo">
  <img height=120 width=300 src="./imgs/logo-white.svg#gh-dark-mode-only" alt="qwantify logo">
</h3>
<p align="center">
  <p align="center">Self hosted cloud gaming with friends, right from the browser</p>
</p>
<h4 align="center">
  <a href="https://qwantify.vercel.app/">Arcade</a> |
  <a href="https://docs.google.com/forms/d/e/1FAIpQLSfBZDOlcdnvJwUdt5ju-v8Gx4oIqHd_jHu_p6QCvlwgaOvQ0A/viewform">Hosting</a> |
  <a href="https://qwantify.vercel.app/">Docs</a> |
  <a href="https://qwantify.vercel.app/">Website</a>
</h4>

<h4 align="center">
    <!-- Not ready for discord yet -->
  <a href="https://github.com/wanjohiryan/qwantify/releases">
      <img src="https://img.shields.io/github/v/release/wanjohiryan/qwantify" alt="release">
    </a>
  <a href="https://github.com/wanjohiryan/qwantify">
    <img src="https://img.shields.io/badge/PRs-Welcome-brightgreen" alt="PRs welcome!" />
  </a>
  <a href="">
    <img src="https://img.shields.io/github/commit-activity/m/wanjohiryan/qwantify" alt="git commit activity" />
  </a>
</h4>

<pre align="center" style="width:100%;padding:40px;">
  <img src="./imgs/play.gif" width="95%" alt="playing with qwantify"/>
</pre>

**[qwantify](https://qwantify.vercel.app)** lets you run PC games (.exe) on a shared host computer with at least one gpu with no extra configurations. _Everything just runs perfectly._

>Note: this was previously a fork of m1k1o's n.eko as a proof of concept. However, as of v0.1.1, they are no  longer backwards compatible.
>This project is (still) in active maintenance. Currently working on the qwantify-core 'engine' [here](https://github.com/wanjohiryan/warp) - gamepad support and nvenc/vaapi encoding.

## 🌞 Motivation
I've always wanted to stream games from different devices while playing them on the fly, and occasionally I even wanted to invite others.
 
Although cloud gaming providers offered this, I preferred a self-hosted version so that I could manage and run my own games.

I then discovered [Parsec](https://www.parsec.app), which was fantastic when it functioned but absolutely useless when the network experienced any little instability.The lack of a web interface and the requirement to install native apps only served to magnify the issue.

I came upon m1k1o's [n.eko](https://github.com/m1k1o/neko) and after making some adjustments for my nvidia GPU, it worked!

Now i could play online with anyone, run multiple games on the same machine, save and sync game progress between computers. It was a miracle.

>And that's how qwantify was born :)

## 💘 Features

> qwantify was, to some degree, inspired by Google Stadia

_Long live Linux 💝_

- **Crowd Play** - play games online with your pals, directly from your browser. Turn any game into multiplayer.
- **State Share** - transfer game play progress between devices or to friends
- **Run Android and Mac executables** - Play games from the Google Play store and the Mac App Store.
- **Get low latency 1080p@60fps game streaming to any browser**
- **Automated and manual gamepad mapping**
- **Multi-device support for all your games**
- **Live stream to Youtube and Twitch**
- **Get automated AMD, Intel, and Nvidia GPU performance tweaks**
- **Url Invites** - send url invitations to friends, even on self-hosted qwantify instances for free

>Check out the [projects tab](https://github.com/wanjohiryan/qwantify/projects?query=is%3Aopen) to get a view on what features have already been shipped

And more.

<p align="center" >
  <a href="https://qwantify.vercel.app" target="_blank"><img src="./imgs/sign-up.png" height=120 /> </a>
  <a href="https://qwantify.vercel.app" target="_blank"><img src="./imgs/game-demo.png" height=120 /> </a>
  <a href="https://docs.google.com/forms/d/e/1FAIpQLSfBZDOlcdnvJwUdt5ju-v8Gx4oIqHd_jHu_p6QCvlwgaOvQ0A/viewform" target="_blank"><img src="./imgs/partner-up.png" height=120></a>
</p>

## 🚀 Get started

To quickly get started, pull the image and run it with docker compose (*recommended*)

Requirements:

1. Linux or WSL
   >qwantify doesn't work on windows/Mac as they cannot pass gpus to linux containers
   
2. Latest version of docker and docker compose
3. A machine with GPU: **Nvidia, AMD or Intel**
  
  >For machines with Nvidia GPUs you will need: `nvidia-docker` and
   `nvidia container toolkit` v450.80.02 or higher

```bash
version: "3.8"
services:
  qwantify:
    image: wanjohiryan/qwantify:latest #or ghcr.io/wanjohiryan/qwantify:20.04
    restart: "unless-stopped"
    ports:
      - "8080:8080" #web interface
    volumes:
      - /games:/games #directory with folders containing your game(s)
    shm_size:'5gb' #size of shared memory
    deploy:
      resources:
        reservations:
          devices: #share nvidia gpu (recommended)
            - capabilities: [gpu] 
        limits:
          memory: 5G #depends on the game (recommended is 4)
          cpus: '4' #depends on the game (recommended is 4)
    devices:
      - /dev/dri:/dev/dri #pass in Intel or AMD GPUs
    environment:
      - APPPATH=/path/to/game/folder #folder containing the game
      - APPFILE=/game.exe #game executable file

```

Then run

```bash
docker-compose up -d
```

## 🔥 What's cool about this?

Not only do you stream games with qwantify, you get the best GPU & CPU performance optimisations, all specifically tailored for the game you're playing.

Additionally, you get high quality 1080p@60fps streams to any browser on the same LAN or online.

>We're on a mission to make games more accessible to all, <i>not just gamers with expensive hardware</i>. 

## 🔄 Comparisons with other software
>Note: qwantify is not **JUST** streaming software

[Parsec](https://parsec.app/):Parsec is not open-source. . It only offers the best performance on Windows or Mac hosts, though and does not function in the browser. It also does not come with performance optimizations pre-installed.

[CloudMorph](https://github.com/giongto35/cloud-morph): Cloudmorph uses WebRTC as opposed to qwantify, which uses QUIC/HTTP3. Additionally, it doesn't implement any hardware acceleration.

[n.eko](https://github.com/m1k1o/neko): neko uses WebRTC as opposed to qwantify, which uses QUIC/HTTP3. It also does not support gamepads/joysticks.

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
         <td>Offline ❗</td>
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
    <tr>
        <th scope="row">Australia and New Zealand</th>
        <td>Currently not available :( ❌</td>
    </tr>
</table>

### Rent out your GPU and help us deliver games to everyone

<p  align="center">
  <a href="https://docs.google.com/forms/d/e/1FAIpQLSfBZDOlcdnvJwUdt5ju-v8Gx4oIqHd_jHu_p6QCvlwgaOvQ0A/viewform" target="_blank"><img src="./imgs/partner-up.png" height=120></a>
</p>



>**Stay frosty :)**
