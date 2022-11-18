ARG UBUNTU_RELEASE=20.04
ARG CUDA_VERSION=11.2.2
#
# STAGE 1: SERVER
#
FROM golang:1.18-bullseye as server
WORKDIR /src

#
# install dependencies
RUN set -eux; apt-get update; \
    apt-get install -y --no-install-recommends git cmake make libx11-dev libxrandr-dev libxtst-dev \
    libgstreamer1.0-dev libgstreamer-plugins-base1.0-dev gstreamer1.0-plugins-base gstreamer1.0-plugins-good gstreamer1.0-plugins-bad gstreamer1.0-plugins-ugly; \
    #
    # install libclipboard
    set -eux; \
    cd /tmp; \
    git clone --depth=1 https://github.com/jtanx/libclipboard; \
    cd libclipboard; \
    cmake .; \
    make -j4; \
    make install; \
    rm -rf /tmp/libclipboard; \
    #
    # clean up
    apt-get clean -y; \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/*

#
# build server
COPY server/ .
RUN go get -v -t -d . && go build -o bin/neko cmd/neko/main.go

#
# STAGE 2: CLIENT
#
FROM node:14-bullseye-slim as client
WORKDIR /src

#
# install dependencies
COPY client/package*.json ./
RUN npm install

#
# build client
COPY client/ .
RUN npm run build

#
# STAGE 3: RUNTIME
#
# FROM debian:bullseye-slim
FROM nvcr.io/nvidia/cudagl:11.2.2-runtime-ubuntu20.04

ARG UBUNTU_RELEASE
ARG CUDA_VERSION
# Make all NVIDIA GPUs visible by default
ARG NVIDIA_VISIBLE_DEVICES=all
# All NVIDIA driver capabilities should preferably be used, check `NVIDIA_DRIVER_CAPABILITIES` inside the container if things do not work
ENV NVIDIA_DRIVER_CAPABILITIES all
ENV VGL_DISPLAY egl
ENV REFRESH 60
#
# avoid warnings by switching to noninteractive
ENV DEBIAN_FRONTEND=noninteractive

#
# set custom user
ARG USERNAME=neko
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN set -eux; \
    apt-get update; \
    #
    # install dependencies
    apt-get install -y --no-install-recommends wget ca-certificates supervisor; \
    apt-get install -y --no-install-recommends pulseaudio dbus-x11 xserver-xorg-video-dummy; \
    apt-get install -y --no-install-recommends libcairo2 libxcb1 libxrandr2 libxv1 libopus0 libvpx6; \
    apt-get install -y gnupg2 software-properties-common cabextract aptitude vim curl; \
    #
    # intel driver + vaapi
    apt-get install -y --no-install-recommends intel-media-va-driver-non-free libva2 vainfo; \
    #
    # gst + vaapi plugin
    apt-get install -y --no-install-recommends libgstreamer1.0-dev libgstreamer-plugins-base1.0-dev \
                    gstreamer1.0-plugins-base gstreamer1.0-plugins-good gstreamer1.0-plugins-bad gstreamer1.0-plugins-ugly gstreamer1.0-pulseaudio \
                    gstreamer1.0-vaapi ;\
    #
    # install fonts
    apt-get install -y --no-install-recommends \
        # Google emojis
        fonts-noto-color-emoji \
        # Japanese fonts
        fonts-takao-mincho \
        # Chinese fonts
        fonts-wqy-zenhei; \
    #
    # create a non-root user
    groupadd --gid $USER_GID $USERNAME; \
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
    # make directories for neko
    mkdir -p /etc/neko /var/www /var/log/neko; \
    chmod 1777 /var/log/neko; \
    chown $USERNAME /var/log/neko/; \
    chown -R $USERNAME:$USERNAME /home/$USERNAME; \
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
# Install wine
ARG WINE_BRANCH="stable"
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

#
# Install graphic libraries //JohnWick uses dx11
RUN winetricks d3dx11_43 
# RUN winetricks --force -q dotnet48

#
# Silence all fixme warnings from wine
ENV WINEDEBUG fixme-all

#
#volume for mounting games
VOLUME [ "/games" ]

#
# copy config files
COPY gamevm/base/dbus /usr/bin/dbus
COPY gamevm/base/default.pa /etc/pulse/default.pa
COPY gamevm/base/supervisord.conf /etc/neko/supervisord.conf
COPY gamevm/base/xorg.conf /etc/neko/xorg.conf
COPY gamevm/base/add-render-group.sh /usr/bin/add-render-group.sh
COPY gamevm/base/entrypoint.sh /etc/entrypoint.sh

#
# set default envs
ENV USER=$USERNAME
ENV DISPLAY=:99.0
ENV NEKO_PASSWORD=neko
ENV NEKO_PASSWORD_ADMIN=admin
ENV NEKO_BIND=:8080
ENV RENDER_GID=

#
# copy static files from previous stages
COPY --from=server /src/bin/neko /usr/bin/neko
COPY --from=client /src/dist/ /var/www

HEALTHCHECK --interval=10s --timeout=5s --retries=8 \
    CMD wget -O - http://localhost:${NEKO_BIND#*:}/health || exit 1

#
# run neko
CMD ["/usr/bin/supervisord", "-c", "/etc/neko/supervisord.conf"]
