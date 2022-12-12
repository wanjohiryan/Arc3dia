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
- **Cloud and GPU Agnostic deployment** that lets you play and host games anywhere anytime, through the browser 
- **Complete control over your game data** - save your game progress locally
- **Play with multiple gamepads** per gameroom. Turn any game into multiplayer
- 🔜 **1-Click Deploy** locally, AWS or GCP
- 🔜 **Url Sharing** for gamerooms (gamepad control switching soon after)
- 🔜 **Official Support for AMD and Intel Gpus**
- 🔜 **Url invites for friends**
- 🔜 **No extra installations needed**
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

- [x] Public Alpha: Anyone can sign up over at [qwantify arcade](https://qwantify.vercel.app/) 
- [ ] Public Beta: Stable enough for most gamers.
- [ ] Public: Production-ready.

We're currently in Public Alpha.
## 🔌 Integrations

We're currently setting the foundation and building a gaming network so games can be played from anywhere on the planet. Any help is welcome! :)

<table>
<tr>
  <th>Platforms </th>
  <th>Frameworks</th>
</tr>
<tr> 
  <td>

<table>
  <tbody>
    <tr>
      <td align="left" valign="middle">
        <a href="https://infisical.com/docs/integrations/platforms/docker?ref=github.com">
          ✔️ Docker
        </a>
      </td>
      <td align="left" valign="middle">
        <a href="https://infisical.com/docs/integrations/platforms/docker-compose?ref=github.com">
          ✔️ Docker Compose
        </a>
      </td>
      <td align="left" valign="middle">
        <a href="https://infisical.com/docs/integrations/cloud/heroku?ref=github.com">
          ✔️ Heroku
        </a>
      </td>
    </tr>
    <tr>
      <td align="left" valign="middle">
        🔜 Vercel (https://github.com/Infisical/infisical/issues/60)
      </td>
      <td align="left" valign="middle">
        🔜 GitLab CI/CD
      </td>
      <td align="left" valign="middle">
        🔜 Fly.io
      </td>
    </tr>
    <tr>
      <td align="left" valign="middle">
        🔜 AWS
      </td>
      <td align="left" valign="middle">
        🔜 GitHub Actions (https://github.com/Infisical/infisical/issues/54)
      </td>
      <td align="left" valign="middle">
         🔜 Railway
      </td>
    </tr>
    <tr>
      <td align="left" valign="middle">
        🔜 GCP
      </td>
      <td align="left" valign="middle">
        🔜 Kubernetes
      </td>
      <td align="left" valign="middle">
        🔜 CircleCI
      </td>
    </tr>
    <tr>
      <td align="left" valign="middle">
        🔜 Jenkins
      </td>
      <td align="left" valign="middle">
        🔜 Digital Ocean
      </td>
      <td align="left" valign="middle">
        🔜 Azure
      </td>
    </tr>
    <tr>
      <td align="left" valign="middle">
         🔜 TravisCI
      </td>
      <td align="left" valign="middle">
         🔜 Netlify (https://github.com/Infisical/infisical/issues/55)
      </td>
    </tr>
  </tbody>
</table>

  </td>
<td>


<table>
  <tbody>
    <tr>
      <td align="left" valign="middle">
        <a href="https://infisical.com/docs/integrations/frameworks/react?ref=github.com">
          ✔️ React
        </a>
      </td>
      <td align="left" valign="middle">
        <a href="https://infisical.com/docs/integrations/frameworks/express?ref=github.com">
          ✔️ Express
        </a>
      </td>
    </tr>
    <tr>
      <td align="left" valign="middle">
        <a href="https://infisical.com/docs/integrations/frameworks/gatsby?ref=github.com">
          ✔️ Gatsby
        </a>
      </td>
      <td align="left" valign="middle">
        <a href="https://infisical.com/docs/integrations/frameworks/flask?ref=github.com">
          ✔️ Flask
        </a>
      </td>
    </tr>
    <tr>
      <td align="left" valign="middle">
        <a href="https://infisical.com/docs/integrations/frameworks/django?ref=github.com">
          ✔️ Django
        </a>
      </td>
      <td align="left" valign="middle">
        <a href="https://infisical.com/docs/integrations/frameworks/laravel?ref=github.com">
          ✔️ Laravel
        </a>
      </td>
    </tr>
    <tr>
      <td align="left" valign="middle">
        <a href="https://infisical.com/docs/integrations/frameworks/nestjs?ref=github.com">
          ✔️ NestJS
        </a>
      </td>
      <td align="left" valign="middle">
        <a href="https://infisical.com/docs/integrations/frameworks/remix?ref=github.com">
          ✔️ Remix
        </a>
      </td>
    </tr>
    <tr>
      <td align="left" valign="middle">
        <a href="https://infisical.com/docs/integrations/frameworks/nextjs?ref=github.com">
          ✔️ Next.js
        </a>
      </td>
      <td align="left" valign="middle">
        <a href="https://infisical.com/docs/integrations/frameworks/vite?ref=github.com">
          ✔️ Vite
        </a>
      </td>
    </tr>
    <tr>
      <td align="left" valign="middle">
        <a href="https://infisical.com/docs/integrations/frameworks/rails?ref=github.com">
          ✔️ Ruby on Rails
        </a>
      </td>
      <td align="left" valign="middle">
        <a href="https://infisical.com/docs/integrations/frameworks/vue?ref=github.com">
          ✔️ Vue
        </a>
      </td>
    </tr>
  </tbody>
</table>

</td>
</tr> 
</table>


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
