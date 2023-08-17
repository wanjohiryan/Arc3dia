#! Some code was copied or inspired by https://github.com/ehfd/nvidia-dind/blob/main/Dockerfile
# Which is licensed under  MPL-2.0 license
# Thanks ehfd :)
FROM ghcr.io/wanjohiryan/ffmpeg/v6:nightly

ARG DEBIAN_FRONTEND=noninteractive

#Accept Steam Terms automatically
SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN echo steam steam/question select "I AGREE" | debconf-set-selections \
 && echo steam steam/license note '' | debconf-set-selections

#
#Install essential libraries
RUN dpkg --add-architecture i386 && apt-get update -y && apt-get install -y --no-install-recommends \
        apt-utils ca-certificates openssh-client vim wget software-properties-common\
        curl iptables git gnupg sudo locales\
        #SteamCMD
        steamcmd; \
        #
        #clean
        rm -rf /var/lib/apt/list/*

#
# Update SteamCMD and verify latest version
RUN locale-gen en_US.UTF-8;\ 
        ln -s /usr/games/steamcmd /usr/bin/steamcmd;\
        steamcmd +quit

#
#Install wine
ARG WINE_BRANCH="staging"
RUN wget -nv -O- https://dl.winehq.org/wine-builds/winehq.key | APT_KEY_DONT_WARN_ON_DANGEROUS_USAGE=1 apt-key add - \
    && apt-add-repository "deb https://dl.winehq.org/wine-builds/ubuntu/ $(grep VERSION_CODENAME= /etc/os-release | cut -d= -f2) main" \
    && dpkg --add-architecture i386 \
    && apt-get update \
    && DEBIAN_FRONTEND="noninteractive" apt-get install -y --install-recommends winehq-${WINE_BRANCH} \
    && rm -rf /var/lib/apt/lists/*

#
#Install gecko and mono installers
COPY docker/nvidia/download_gecko_and_mono.sh /root/download_gecko_and_mono.sh
RUN chmod +x /root/download_gecko_and_mono.sh \
    && /root/download_gecko_and_mono.sh "$(dpkg -s wine-${WINE_BRANCH} | grep "^Version:\s" | awk '{print $2}' | sed -E 's/~.*$//')"

#
# Install winetricks
RUN wget -nv -O /usr/bin/winetricks https://raw.githubusercontent.com/Winetricks/winetricks/master/src/winetricks \
    && chmod +x /usr/bin/winetricks

#
# NVIDIA Container Toolkit & Docker
RUN distribution=$(. /etc/os-release;echo $ID$VERSION_ID); \
    curl -s -L https://nvidia.github.io/nvidia-docker/gpgkey | apt-key add - ; \
    curl -s -L https://nvidia.github.io/nvidia-docker/$distribution/nvidia-docker.list | tee /etc/apt/sources.list.d/nvidia-docker.list; \
    apt-get update && apt-get install -y nvidia-docker2 docker.io docker-compose; \
    #
    #clean
    rm -rf /var/lib/apt/list/*

ARG USERNAME=arc3dia
ARG USER_UID=1000
ARG USER_GID=${USER_UID}

#
# Add non-root user
RUN groupadd --gid ${USER_GID} ${USERNAME}; \
        useradd --uid ${USER_UID} --gid ${USERNAME} --shell /bin/bash --create-home ${USERNAME};\
        usermod -a -G sudo ${USERNAME} ;\
        echo "${USERNAME} ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers; \
        #
        #Make sure the home directory is owned by user
        chown -R $USERNAME:$USERNAME /home/${USERNAME}

#
#Copy files
COPY docker/nvidia/modprobe docker/nvidia/ffmpeg.sh docker/nvidia/entrypoint.sh docker/nvidia/dockerd.sh /usr/local/bin/
COPY docker/nvidia/logger.sh /opt/bash-utils/logger.sh
COPY ./arc3dia /usr/local/bin/

#
#Add execution permissions
RUN chmod +x /usr/local/bin/entrypoint.sh /usr/local/bin/dockerd.sh /usr/local/bin/modprobe /usr/local/bin/arc3dia /usr/local/bin/ffmpeg.sh

#
#Install tini https://github.com/krallin/tini
ARG TINI_VERSION=v0.19.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /usr/bin/tini
RUN chmod +x /usr/bin/tini
ENTRYPOINT [ "/usr/bin/tini","--", "/usr/local/bin/dockerd.sh" ]

#
#Add env variables
ENV USER ${USERNAME}
ENV SHELL /bin/bash
ENV LANG 'en_US.UTF-8'
ENV LANGUAGE 'en_US:en'
#
#volume for mounting 
#1.preinstalled game #2. config file #3. game data for persisting games
#directories /certs and /media are for testing only
VOLUME [ "/game", "/arc3dia", "/data", "/certs", "/media" ]

#
#Run the bash shell as a non-root user
USER ${USERNAME}
WORKDIR /home/${USERNAME}

CMD ["/usr/local/bin/entrypoint.sh"]