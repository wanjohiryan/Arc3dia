ARG BASE_IMAGE=m1k1o/neko:base
FROM $BASE_IMAGE

#
#Install essentials
RUN apt update \
    && apt-get update -y \
    && apt-get install -y gnupg\
    && apt-get install --no-install-recommends --assume-yes wget software-properties-common gpg-agent mingw-w64 cabextract aptitude vim\
    && apt-get clean \
    && apt-get autoremove

#
# Install wine
ARG WINE_BRANCH="stable"
RUN dpkg --add-architecture i386 \
    && mkdir -pm755 /etc/apt/keyrings \
    && wget -O /etc/apt/keyrings/winehq-archive.key https://dl.winehq.org/wine-builds/winehq.key \
    && wget -NP /etc/apt/sources.list.d/ https://dl.winehq.org/wine-builds/debian/dists/bullseye/winehq-bullseye.sources \
    && apt update \
    && DEBIAN_FRONTEND="noninteractive" apt-get install --install-recommends winehq-${WINE_BRANCH} \
    # && rm -rf /var/lib/apt/lists/*

#
# Install winetricks
RUN wget -O /usr/bin/winetricks https://raw.githubusercontent.com/Winetricks/winetricks/master/src/winetricks \
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
