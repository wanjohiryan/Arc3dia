<h2 height="110 align="center">
  <img width="500" height= src="/imgs/logo-black.svg#gh-light-mode-only" alt="qwantify logo">
  <img width="500" src="/imgs/logo-white.svg#gh-dark-mode-only" alt="qwantify logo">
</h2>
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

<img src="/imgs/signup-black.png#gh-light-mode-only" width="100%" alt="qwantify signup" />
<img src="/imgs/signup-white.png#gh-dark-mode-only" width="100%"  alt="qwantify signup" />



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

```bash
docker-compose up -d docker-compose.yaml
```

<p>
  <a href="https://infisical.com/docs/self-hosting/overview" target="_blank"><img src="https://user-images.githubusercontent.com/78047717/206356882-2b773eed-b0da-4725-ae2f-83e3cd7f2713.png" height=120 /> </a>
  <a href="https://www.youtube.com/watch?v=JS3OKYU2078" target="_blank"><img src="https://user-images.githubusercontent.com/78047717/206356600-8833b128-6cae-408c-a703-07b2fc6aff4b.png" height=120 /> </a>
  <a href="https://app.infisical.com/signup" target="_blank"><img src="https://user-images.githubusercontent.com/78047717/206355970-f4c09062-b88f-452a-94e0-9c61a0651170.png" height=120></a>
</p>

## 🔥 What's cool about this?

qwantify makes gaming . We're on a mission to make games more accessible to all, <i>not just gamers with expensive hardware</i>. 

We are currently working hard to make qwantify more extensive. Need any integrations or want a new feature? Feel free to [create an issue](https://github.com/wanjohiryan/qwantify/issues) or [contribute](https://infisical.com/docs/contributing/overview) directly to the repository.

## 🌱 Contributing

Whether it's big or small, we love contributions ❤️ Check out our guide to see how to [get started](https://infisical.com/docs/contributing/overview).

Not sure where to get started? You can:
- [Book a free, non-pressure pairing sessions with one of our teammates](mailto:tony@infisical.com?subject=Pairing%20session&body=I'd%20like%20to%20do%20a%20pairing%20session!)!
- Join our <a href="https://join.slack.com/t/infisical-users/shared_invite/zt-1kdbk07ro-RtoyEt_9E~fyzGo_xQYP6g">Slack</a>, and ask us any questions there.

## 💚 Community & Support

- [Slack](https://join.slack.com/t/infisical-users/shared_invite/zt-1kdbk07ro-RtoyEt_9E~fyzGo_xQYP6g) (For live discussion with the community and the Infisical team)
- [GitHub Discussions](https://github.com/Infisical/infisical/discussions) (For help with building and deeper conversations about features)
- [GitHub Issues](https://github.com/Infisical/infisical-cli/issues) (For any bugs and errors you encounter using Infisical)
- [Twitter](https://twitter.com/infisical) (Get news fast) 

## 🐥 Status

- [x] Public Alpha: Anyone can sign up over at [infisical.com](https://infisical.com) but go easy on us, there are kinks and we're just getting started.
- [ ] Public Beta: Stable enough for most non-enterprise use-cases.
- [ ] Public: Production-ready.

We're currently in Public Alpha.

## 🚨 Stay Up-to-Date

Infisical officially launched as v.1.0 on November 21st, 2022. However, a lot of new features are coming very quickly. Watch **releases** of this repository to be notified about future updates:

![infisical-star-github](https://github.com/Infisical/infisical/blob/main/.github/images/star-infisical.gif?raw=true)

## 🔌 Integrations

We're currently setting the foundation and building [integrations](https://infisical.com/docs/integrations/overview) so secrets can be synced everywhere. Any help is welcome! :)

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
