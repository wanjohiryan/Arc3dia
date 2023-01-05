ARG UBUNTU_RELEASE=20.04
ARG CUDA_VERSION=11.2.2
FROM nvcr.io/nvidia/cudagl:${CUDA_VERSION}-runtime-ubuntu${UBUNTU_RELEASE}

#
## for ghcr
LABEL maintainer "https://github.com/wanjohiryan"

# Make all NVIDIA GPUs visible by default
ARG NVIDIA_VISIBLE_DEVICES=all
# All NVIDIA driver capabilities should preferably be used, check `NVIDIA_DRIVER_CAPABILITIES` inside the container if things do not work
ENV NVIDIA_DRIVER_CAPABILITIES all
#
# avoid warnings by switching to noninteractive
ENV DEBIAN_FRONTEND=noninteractive
#
# set vgl-display to headless 3d gpu card
# correct values are egl[n] or /dev/dri/card0:if this is passed into container
ENV VGL_DISPLAY egl
#
#for vulkan //creates complications
ENV XDG_RUNTIME_DIR /tmp
#
# set custom user
ARG USERNAME=qwantify
ARG USER_UID=1000
ARG USER_GID=$USER_UID
ENV USER=$USERNAME

#
#Install go v1.18 //quic-go does not work on v1.19 (as of 1-1-23)
ARG GO_VERSION=1.18.9
ENV GO_VERSION=${GO_VERSION}

#
# Install locales to prevent Xvfb errors
RUN apt-get clean && \
    apt-get update && apt-get install --no-install-recommends -y locales && \
    rm -rf /var/lib/apt/lists/* && \
    locale-gen en_US.UTF-8
ENV LANG en_US.UTF-8
ENV LANGUAGE en_US:en
ENV LC_ALL en_US.UTF-8

#
# install dependencies: xvfb pulseaudio dbus wget curl supervisoretc
RUN set -eux; \
    apt-get update; \
    apt-get install -y --no-install-recommends wget ca-certificates supervisor; \
    apt-get install -y --no-install-recommends pulseaudio dbus-x11 xvfb mingw-w64 ffmpeg; \
    apt-get install -y --no-install-recommends libcairo2 libxcb1 libxrandr2 libxv1 libopus0 libvpx6; \
    apt-get install -y gnupg2 software-properties-common cabextract aptitude vim curl;\
    #
    # clean up
    apt-get clean -y; \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/*

RUN apt-get update; \
    apt-get install -y git gcc; \
    wget -P /tmp "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz"; \
    tar -C /usr/local -xzf "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"; \
    rm "/tmp/go${GO_VERSION}.linux-amd64.tar.gz";

ENV GOPATH /go

ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

#
#Install nodejs
RUN apt-get update; \
    apt-get install curl  -y; \
    curl -sL https://deb.nodesource.com/setup_18.x | bash; \
    apt-get install nodejs -y;

#
# xvfb display
ENV DISPLAY=:99.0

#
# mkcert install
RUN apt-get update; \
    apt install libnss3-tools -y;\
    wget -nv -O /usr/local/bin/mkcert https://github.com/FiloSottile/mkcert/releases/download/v1.4.3/mkcert-v1.4.3-linux-amd64; \
    chmod +x /usr/local/bin/mkcert;

# 
#On second thought, let the user change this from the frontend rather than on startup
#TODO: implement this on from the webpage
# #1920x1080@60fps default
# ENV SCREEN_WIDTH 1920
# ENV SCREEN_HEIGHT 1080
# ENV VIDEO_FPS 60

#TODO: implement this
#video encoder: valid values are: nvenc(nvidia), vaapi(amd/intel) or x264(software-rendering)#must have 6cores or more
ENV VIDEO_ENCODER x264

#
## hardware encoding for /dev/dri
ENV RENDER_GID=

#
# hardware acclerations utilities
RUN set -eux; \
    apt-get update; \
    apt-get install --no-install-recommends -y libgtk-3-bin mesa-utils mesa-utils-extra mesa-va-drivers mesa-vulkan-drivers libvulkan-dev libvulkan-dev:i386 vdpauinfo; \
    #
    # intel driver + vaapi
    apt-get install -y --no-install-recommends intel-media-va-driver-non-free libva2 vainfo;\
    #
    # clean up
    apt-get clean -y; \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/*

#
# create a non-root user
RUN groupadd --gid $USER_GID $USERNAME; \
    useradd --uid $USER_UID --gid $USERNAME --shell /bin/bash --create-home $USERNAME; \
    adduser $USERNAME audio; \
    adduser $USERNAME video; \
    adduser $USERNAME pulse; \
    #
    # setup pulseaudio
    mkdir -p /home/$USERNAME/.config/pulse/; \
    echo "default-server=unix:/tmp/pulseaudio.socket" > /home/$USERNAME/.config/pulse/client.conf; \
    #
    # workaround for an X11 problem: http://blog.tigerteufel.de/?p=476
    mkdir /tmp/.X11-unix; \
    chmod 1777 /tmp/.X11-unix; \
    chown $USERNAME /tmp/.X11-unix/; \
    #
    # clean up
    apt-get clean -y; \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/*

#
# Install and configure Vulkan manually
RUN if [ "${UBUNTU_RELEASE}" = "18.04" ]; then apt-get update && apt-get install --no-install-recommends -y vulkan-utils; else apt-get update && apt-get install --no-install-recommends -y vulkan-tools; fi && \
    rm -rf /var/lib/apt/lists/* && \
    VULKAN_API_VERSION=$(dpkg -s libvulkan1 | grep -oP 'Version: [0-9|\.]+' | grep -oP '[0-9]+(\.[0-9]+)(\.[0-9]+)') && \
    mkdir -p /etc/vulkan/icd.d/ && \
    echo "{\n\
    \"file_format_version\" : \"1.0.0\",\n\
    \"ICD\": {\n\
        \"library_path\": \"libGLX_nvidia.so.0\",\n\
        \"api_version\" : \"${VULKAN_API_VERSION}\"\n\
    }\n\
}" > /etc/vulkan/icd.d/nvidia_icd.json

ARG VIRTUALGL_VERSION=3.0.2
#
# Install VirtualGL and make libraries available for preload
RUN curl -fsSL -O "https://sourceforge.net/projects/virtualgl/files/virtualgl_${VIRTUALGL_VERSION}_amd64.deb" && \
    curl -fsSL -O "https://sourceforge.net/projects/virtualgl/files/virtualgl32_${VIRTUALGL_VERSION}_amd64.deb" && \
    apt-get update && apt-get install -y --no-install-recommends ./virtualgl_${VIRTUALGL_VERSION}_amd64.deb ./virtualgl32_${VIRTUALGL_VERSION}_amd64.deb && \
    rm -f "virtualgl_${VIRTUALGL_VERSION}_amd64.deb" "virtualgl32_${VIRTUALGL_VERSION}_amd64.deb" && \
    rm -rf /var/lib/apt/lists/* && \
    chmod u+s /usr/lib/libvglfaker.so && \
    chmod u+s /usr/lib/libdlfaker.so && \
    chmod u+s /usr/lib32/libvglfaker.so && \
    chmod u+s /usr/lib32/libdlfaker.so && \
    chmod u+s /usr/lib/i386-linux-gnu/libvglfaker.so && \
    chmod u+s /usr/lib/i386-linux-gnu/libdlfaker.so

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
COPY gamevm/base/download_gecko_and_mono.sh /root/download_gecko_and_mono.sh
RUN chmod +x /root/download_gecko_and_mono.sh \
    && /root/download_gecko_and_mono.sh "$(dpkg -s wine-${WINE_BRANCH} | grep "^Version:\s" | awk '{print $2}' | sed -E 's/~.*$//')"

#TODO: replace with dxvk
#install dx11
RUN winetricks d3dx11_43

#
# Silence all fixme warnings from wine
ENV WINEDEBUG fixme-all

WORKDIR /qwantify

#
# Copy client and install node-modules
COPY internal/client/ ./client

RUN cd ./client && npm install
#
# Copy cert folder and make certs
COPY internal/cert ./cert

RUN cd ./cert && chmod +x ./generate-cert.sh && ./generate-cert.sh

#
# Make a directory to store files from ffmpeg
RUN mkdir ./media

#
# Copy server
COPY internal/server ./server

#
#volume for mounting games
VOLUME [ "/games" ]

#
#expose only one port
EXPOSE 8080

#
# copy config files
COPY gamevm/base/dbus /usr/bin/dbus
COPY gamevm/base/default.pa /etc/pulse/default.pa
COPY gamevm/base/supervisord.conf /etc/qwantify/supervisord.conf
COPY gamevm/base/add-render-group.sh /usr/bin/add-render-group.sh
COPY gamevm/base/run-the-game.sh /etc/run-the-game.sh
COPY gamevm/base/ffmpeg.sh /etc/ffmpeg.sh

#
# run qwantify using CMD, this way it's easier to ssh into a running container
CMD ["/usr/bin/supervisord", "-c", "/etc/qwantify/supervisord.conf"]
