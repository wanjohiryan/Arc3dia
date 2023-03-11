#! Some code was copied or inspired by https://github.com/selkies-project/docker-nvidia-glx-desktop
# Which is licensed under  MPL-2.0 license
# Thanks ehfd :)

#
#Import base image from warp
# Latest stable major version is 0 but nvenc is at nightly for now
# Comes with vulkan and egl installed
#
FROM ghcr.io/wanjohiryan/warp/nvenc:latest

#
#Use the root user from here
#
USER root


#
# Install wine >= 7.2 for vk3d capabilities // probably not that stable :)
ARG WINE_BRANCH="staging"
RUN wget -nv -O- https://dl.winehq.org/wine-builds/winehq.key | APT_KEY_DONT_WARN_ON_DANGEROUS_USAGE=1 apt-key add - \
    && apt-add-repository "deb https://dl.winehq.org/wine-builds/ubuntu/ $(grep VERSION_CODENAME= /etc/os-release | cut -d= -f2) main" \
    && dpkg --add-architecture i386 \
    && apt-get update \
    && DEBIAN_FRONTEND="noninteractive" apt-get install -y --install-recommends winehq-${WINE_BRANCH} \
    && rm -rf /var/lib/apt/lists/*

#
# Install winetricks
RUN wget -nv -O /usr/bin/winetricks https://raw.githubusercontent.com/Winetricks/winetricks/master/src/winetricks \
    && chmod +x /usr/bin/winetricks

#
# Download gecko and mono installers
COPY docker/nvidia/download_gecko_and_mono.sh /root/download_gecko_and_mono.sh
RUN chmod +x /root/download_gecko_and_mono.sh \
    && /root/download_gecko_and_mono.sh "$(dpkg -s wine-${WINE_BRANCH} | grep "^Version:\s" | awk '{print $2}' | sed -E 's/~.*$//')"

#
#install dx11 and dx9
#! Use dxvk instead, winetricks adds >600mb  
# RUN winetricks -q d3dx9  d3dx11

#
# Silence all fixme warnings from wine
ENV WINEDEBUG fixme-all

#
#volume for mounting games
VOLUME [ "/games" ]

#
#Copy entrypoint and movr it to where warp can see
COPY docker/nvidia/qwantify.sh /etc/warp/entrypoint.d/qwantify.sh
RUN chmod +x /etc/warp/entrypoint.d/qwantify.sh

#
#return back to our designated user
USER ${USERNAME}