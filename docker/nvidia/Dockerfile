#! Some code was copied or inspired by https://github.com/ehfd/nvidia-dind/blob/main/Dockerfile
# Which is licensed under  MPL-2.0 license
# Thanks ehfd :)
ARG UBUNTU_RELEASE=22.04
FROM ubuntu:${UBUNTU_RELEASE}

ARG DEBIAN_FRONTEND=noninteractive

#Accept Steam Terms automatically
SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN echo steam steam/question select "I AGREE" | debconf-set-selections \
 && echo steam steam/license note '' | debconf-set-selections

#
#Install essential libraries
RUN apt-get update && dpkg --add-architecture i386 && apt-get install -y \
        apt-utils ca-certificates openssh-client vim wget \
        curl iptables git gnupg sudo locales\
        #SteamCMD
        steamcmd; \
        #
        #clean
        rm -rf /var/lib/apt/list/*;\
        #
        # Update SteamCMD and verify latest version
        locale-gen en_US.UTF-8 && steamcmd +quit

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
COPY docker/nvidia/modprobe docker/nvidia/entrypoint.sh /usr/local/bin/
COPY docker/nvidia/logger.sh /opt/bash-utils/logger.sh

#
#Add execution permissions
RUN chmod +x /usr/local/bin/entrypoint.sh /usr/local/bin/modprobe

#
#Install tini https://github.com/krallin/tini
ARG TINI_VERSION=v0.19.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /usr/bin/tini
RUN chmod +x /usr/bin/tini
ENTRYPOINT [ "/usr/bin/tini","--", "/usr/local/bin/entrypoint.sh" ]

#
#Add env variables
ENV USER ${USERNAME}
ENV SHELL /bin/bash
ENV LANG 'en_US.UTF-8'
ENV LANGUAGE 'en_US:en'
#
#volume for mounting 
#1.preinstalled game #2. config file #3. game data for persisting games
VOLUME [ "/game", "/arc3dia", "/data" ]

#
#Run the bash shell as a non-root user
USER ${USERNAME}
WORKDIR /home/${USERNAME}

CMD ["sh"]