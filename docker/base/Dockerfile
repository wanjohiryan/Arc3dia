# Usage:
#
# $ docker run --runtime=sysbox-runc -it --rm --name=syscont ghcr.io/wanjohiryan/arc3dia:base
#
# This will run systemd, sshd and docker;
#https://github.com/nestybox/dockerfiles/blob/master/ubuntu-jammy-systemd/Dockerfile
FROM ghcr.io/wanjohiryan/ffmpeg/v6:nightly
#
# Systemd installation
#
RUN apt-get update; \
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
RUN apt-get clean && \
    apt-get update && apt-get install --no-install-recommends -y locales && \
    rm -rf /var/lib/apt/lists/* && \
    locale-gen en_US.UTF-8
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
RUN apt-get update; \
    apt-get upgrade; \
    apt-get install -y --no-install-recommends \
    #
    #Xserver 
    wget vim ca-certificates software-properties-common gpg-agent xvfb dbus-user-session dbus-x11 libdbus-c++-1-0v5\
    #
    #Pulseaudio
    pulseaudio libcairo2 libxcb1 libxrandr2 libxv1 libopus0 libvpx6; \
    #
    # workaround for an X11 problem: http://blog.tigerteufel.de/?p=476
    #
    mkdir /tmp/.X11-unix; \
    chmod 1777 /tmp/.X11-unix; \
    chown $USERNAME /tmp/.X11-unix/; \
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
#Create a non-root user `admin`
#
RUN apt-get update; \
    #
    #create a non-root user
    #
    groupadd --gid ${USER_GID} ${USERNAME}; \
    useradd --uid ${USER_UID} --gid ${USERNAME} --shell /bin/bash --create-home ${USERNAME};\
    usermod -a -G adm,audio,input,pulse-access,sudo,tape,tty,video,ffmpeg ${USERNAME} ;\
    echo "${USERNAME} ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers; \
    #
    #make directories for arc3dia
    #
    mkdir -p /certs /scripts;\
    chmod 1777 /certs /scripts; \
    chown ${USERNAME} /certs /scripts; \
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
RUN apt-get update && apt-get install -y curl; \
    rm -rf /var/lib/apt/lists/*; \
    curl -fsSL https://get.docker.com -o get-docker.sh && sh get-docker.sh; \
    #
    # Add user "admin" to the Docker group
    #
    usermod -a -G docker admin

ADD https://raw.githubusercontent.com/docker/docker-ce/master/components/cli/contrib/completion/bash/docker /etc/bash_completion.d/docker.sh

#
# Install sshd
#
RUN apt-get update && apt-get install --no-install-recommends -y openssh-server \
    && rm -rf /var/lib/apt/lists/* \
    && mkdir /home/admin/.ssh \
    && chown admin:admin /home/admin/.ssh

#
# Make use of stopsignal (instead of sigterm) to stop systemd containers.
#
STOPSIGNAL SIGRTMIN+3

#
#Copy Warp [CLIENT] linux binary and the start script into the container
#
COPY temp/warp /usr/bin/
COPY docker/base/ffmpeg.sh /scripts/
RUN chmod +x /scripts/ffmpeg.sh /usr/bin/warp

COPY docker/base/dbus /usr/bin/dbus
COPY docker/base/default.pa /etc/pulse/

#
#Copy the systemd.service file for starting the server
#
COPY docker/base/*.service /lib/systemd/system/
RUN ln -sf /lib/systemd/system/warp.service \
       /etc/systemd/system/multi-user.target.wants/warp.service
RUN ln -sf /lib/systemd/system/xserver.service \
       /etc/systemd/system/multi-user.target.wants/xserver.service
RUN ln -sf /lib/systemd/system/dbus.service \
       /etc/systemd/system/multi-user.target.wants/dbus.service
RUN ln -sf /lib/systemd/system/pulseaudio.service \
       /etc/systemd/system/multi-user.target.wants/pulseaudio.service
#
# For ssh, and http3 connections respectively
# Warp runs on port 4443 
#

# FIXME: Remove this
#After a lot of thought, i commented out this as running the warp relay server on a local machine might not be the ideal way to do this.
#Ideally we should have the client run locally, transmit the audio and video to a remote warp relay server that will transmit the content to all available clients.

# EXPOSE 22 4443/udp 4443 80

#the volume that houses the key and cert files as `/certs/key.pem` and `/certs/cert.pem`
VOLUME [ "/certs" ]

ENV DISPLAY :0
ENV DPI 96
ENV CDEPTH 24

#
# Set systemd as entrypoint.
#
ENTRYPOINT [ "/sbin/init", "--log-level=err" ]