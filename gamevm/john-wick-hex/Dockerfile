ARG BASE_IMAGE=m1k1o/neko:base
FROM $BASE_IMAGE

#
#Install essentials
RUN apt update \
    && apt-get update -y \
    && apt-get install -y gnupg\
    && apt-get clean \
    && apt-get autoremove

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
COPY download_gecko_and_mono.sh /root/download_gecko_and_mono.sh
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
# copy configuation files
COPY supervisord.conf /etc/neko/supervisord/johnwick.conf
# COPY openbox.xml /etc/neko/openbox.xml
