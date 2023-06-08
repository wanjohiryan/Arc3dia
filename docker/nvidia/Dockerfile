#! Some code was copied or inspired by https://github.com/ehfd/nvidia-dind/blob/main/Dockerfile
# Which is licensed under  MPL-2.0 license
# Thanks ehfd :)
ARG UBUNTU_RELEASE=22.04
FROM ubuntu:${UBUNTU_RELEASE}

#
#Install essential libraries
RUN apt-get update && apt-get install -y \
        apt-utils ca-certificates openssh-client vim\
        curl iptables git gnupg; \
        #
        #clean
        rm -rf /var/lib/apt/list/*

#
# NVIDIA Container Toolkit & Docker
RUN distribution=$(. /etc/os-release;echo $ID$VERSION_ID); \
    curl -s -L https://nvidia.github.io/nvidia-docker/gpgkey | apt-key add - ; \
    curl -s -L https://nvidia.github.io/nvidia-docker/$distribution/nvidia-docker.list | tee /etc/apt/sources.list.d/nvidia-docker.list; \
    apt-get update && apt-get install -y nvidia-docker2 docker.io docker-compose; \
    #
    #clean
    rm -rf /var/lib/apt/list/*

#
#volume for mounting 
#1.preinstalled game #2. config file #3. game data for persisting games
VOLUME [ "/game", "/arc3dia", "/data" ]

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

CMD ["sh"]