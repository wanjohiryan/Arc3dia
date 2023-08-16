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

ARG USERNAME=admin
ARG USER_UID=1000 
ARG USER_GID=${USER_UID}

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
    mkdir -p /certs /media;\
    chmod 1777 /certs /media; \
    chown ${USERNAME} /certs /media; \
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
# Install Sshd
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
# For ssh, and http3 connections respectively
#
EXPOSE 22 443/udp 443 80

#
# Set systemd as entrypoint.
#
ENTRYPOINT [ "/sbin/init", "--log-level=err" ]
