# Usage:
#
# $ docker run --runtime=sysbox-runc -it --rm --name=syscont ghcr.io/wanjohiryan/arc3dia:base
#
# This will run systemd, sshd and docker;
#https://github.com/nestybox/dockerfiles/blob/master/ubuntu-jammy-systemd/Dockerfile
FROM ubuntu:jammy

#
# For selkies
ARG UBUNTU_RELEASE=22.04

#
# Systemd installation
#
RUN apt-get update -y; \
    apt-get install -y --no-install-recommends   \
            systemd \
            systemd-sysv \
            libsystemd0 \
            ca-certificates \
            dbus \
            iptables \
            iproute2 \
            kmod \
            locales \
            sudo \
            udev; \
    #
    # Prevents journald from reading kernel messages from /dev/kmsg
    #
    echo "ReadKMsg=no" >> /etc/systemd/journald.conf; \
    #
    # Housekeeping
    #
    apt-get clean -y; \
    rm -rf \
       /var/cache/debconf/* \
       /var/lib/apt/lists/* \
       /var/log/* \
       /tmp/* \
       /var/tmp/* \
       /usr/share/doc/* \
       /usr/share/man/* \
       /usr/share/local/*

#
# Install locales to prevent X11 errors
#
RUN locale-gen en_US.UTF-8
ENV LANG en_US.UTF-8
ENV LANGUAGE en_US:en
ENV LC_ALL en_US.UTF-8
ENV TZ UTC

ARG USERNAME=admin
ARG USER_UID=1000 
ARG USER_GID=${USER_UID}

#
#Add dependencies
#
RUN apt-get update -y; \
    apt-get upgrade -y; \
    apt-get install -y --no-install-recommends \
    wget \
    vim \
    software-properties-common \
    gpg-agent \
    dbus-user-session \
    dbus-x11 \
    libdbus-c++-1-0v5\
    #
    #Xvfb
    xvfb \
    #
    #Pulseaudio
    pulseaudio; \
    #
    # Housekeeping
    #
    apt-get clean -y; \
    rm -rf \
       /var/cache/debconf/* \
       /var/lib/apt/lists/* \
       /var/log/* \
       /tmp/* \
       /var/tmp/* \
       /usr/share/doc/* \
       /usr/share/man/* \
       /usr/share/local/*

#
#Install ffmpeg v6 this has never been reliable for me. Do imma get swap this for a better alternative later on
#
RUN apt-get update -y; \
    apt-get upgrade -y; \
    add-apt-repository ppa:savoury1/ffmpeg4 \
    add-apt-repository ppa:savoury1/ffmpeg6 \
    apt-get update -y; \
    apt-get upgrade -y && apt-get dist-upgrade -y; \
    apt-get install ffmpeg -y; \
    #
    # Log out the ffmpeg version
    ffmpeg -version


#
#Create a non-root user `admin`
#
RUN apt-get update; \
    #
    #create a non-root user
    #
    groupadd --gid ${USER_GID} ${USERNAME}; \
    useradd --uid ${USER_UID} --gid ${USERNAME} --shell /bin/bash --create-home ${USERNAME};\
    usermod -a -G adm,audio,input,pulse-access,sudo,tape,tty,video,ffmpeg ${USERNAME} ;\
    #
    #Remove sudo password
    echo "${USERNAME} ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers; \
    #
    #make directories for arc3dia
    #
    mkdir -p /certs /scripts /media;\
    chmod 1777 /certs /scripts /media; \
    chown ${USERNAME} /certs /media /scripts; \
    chown -R $USERNAME:$USERNAME /home/${USERNAME} ;\
    ln -snf "/usr/share/zoneinfo/$TZ" /etc/localtime && echo "$TZ" > /etc/timezone

#
# Disable systemd services/units that are unnecessary within a container.
#
RUN systemctl mask systemd-udevd.service \
                   systemd-udevd-kernel.socket \
                   systemd-udevd-control.socket \
                   systemd-modules-load.service \
                   sys-kernel-debug.mount \
                   sys-kernel-tracing.mount

#
# Install Docker
#
RUN apt-get update -y && apt-get install -y curl; \
    rm -rf /var/lib/apt/lists/*; \
    curl -fsSL https://get.docker.com -o get-docker.sh && sh get-docker.sh; \
    #
    # Add user "admin" to the Docker group
    #
    groupadd docker; \
    usermod -a -G docker $USERNAME; \
    newgrp docker

ADD https://raw.githubusercontent.com/docker/docker-ce/master/components/cli/contrib/completion/bash/docker /etc/bash_completion.d/docker.sh

#
# Install sshd
#
RUN apt-get update && apt-get install --no-install-recommends -y openssh-server; \
    rm -rf /var/lib/apt/lists/* ;\
    mkdir /home/${USERNAME}/.ssh; \
    chown $USERNAME:$USERNAME /home/${USERNAME}/.ssh

#
#Install Dan Isla's Js-interposer reference link ---> https://github.com/selkies-project/selkies-gstreamer/tree/main/addons/js-interposer
#
RUN  SELKIES_VERSION=$(curl -fsSL "https://api.github.com/repos/selkies-project/selkies-gstreamer/releases/latest" | jq -r '.tag_name' | sed 's/[^0-9\.\-]*//g') ; \
        curl -O -fsSL "https://github.com/selkies-project/selkies-gstreamer/releases/download/v${SELKIES_VERSION}/selkies-js-interposer-v${SELKIES_VERSION}-ubuntu${UBUNTU_RELEASE}.deb"; \
        apt-get update; \
        apt-get install --no-install-recommends -y "./selkies-js-interposer-v${SELKIES_VERSION}-ubuntu${UBUNTU_RELEASE}.deb"; \
        rm -f "selkies-js-interposer-v${SELKIES_VERSION}-ubuntu${UBUNTU_RELEASE}.deb";\
        rm -rf /var/lib/apt/lists/* 

# Add configuration for Selkies-GStreamer Joystick interposer
ENV LD_PRELOAD /usr/local/lib/selkies-js-interposer/joystick_interposer.so${LD_PRELOAD:+:${LD_PRELOAD}}
ENV SDL_JOYSTICK_DEVICE /dev/input/js0


#
# Make use of stopsignal (instead of sigterm) to stop systemd containers.
#
STOPSIGNAL SIGRTMIN+3

#
#Copy Warp [CLIENT] linux binary and the start script into the container
#
COPY temp/warp /usr/bin/
COPY docker/base/entrypoint.sh /scripts/
RUN chmod +x /scripts/entrypoint.sh /usr/bin/warp

#
#Copy the systemd.service file for starting the server
#
COPY docker/base/*.service /lib/systemd/system/
RUN ln -sf /lib/systemd/system/warp.service \
       /etc/systemd/system/multi-user.target.wants/warp.service
# RUN ln -sf /lib/systemd/system/dbus.service \
#        /etc/systemd/system/multi-user.target.wants/dbus-server.service
RUN ln -sf /lib/systemd/system/pulseaudio.service \
       /etc/systemd/system/multi-user.target.wants/pulseaudio.service
RUN ln -sf /lib/systemd/system/xserver.service \
       /etc/systemd/system/multi-user.target.wants/xserver.service
#
# For ssh, and http3 connections respectively
# Warp runs on port 4443 
#


#RUN ln -snf "/usr/share/zoneinfo/$TZ" /etc/localtime && echo "$TZ" | sudo tee /etc/timezone

# FIXME: Remove this
#After a lot of thought, i commented out this as running the warp relay server on a local machine might not be the ideal way to do this.
#Ideally we should have the client run locally, transmit the audio and video to a remote warp relay server that will transmit the content to all available clients.

# EXPOSE 22 4443/udp 4443 80

#the volume that houses the key and cert files as `/certs/key.pem` and `/certs/cert.pem`
VOLUME [ "/certs" ]

ENV DISPLAY :0
ENV DPI 96
ENV CDEPTH 24
ENV USER ${USERNAME}
ENV SHELL /bin/bash

ENV WARP_SERVER_URL "${WARP_SERVER_URL:-"https://localhost"}"
ENV WARP_SERVER_PORT "${WARP_SERVER_PORT:-4443}"

#Full server url
ENV WARP_SERVER_FULL_URL="${WARP_SERVER_URL}:${WARP_SERVER_PORT}"

#
# Run entrypoint.sh then set systemd as entrypoint.
#
ENTRYPOINT ["/scripts/entrypoint.sh" ]
