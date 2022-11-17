ARG BASE_IMAGE=m1k1o/neko:base
FROM $BASE_IMAGE

#
#Install common files
ARG WINE_BRANCH="stable"
RUN dpkg --add-architecture i386 \
    && apt update \
    && apt -y install gnupg2 software-properties-common cabextract  gpg-agent mingw-w64 cabextract aptitude vim

#
# Install wine (source: https://computingforgeeks.com/how-to-install-wine-on-debian/ )
RUN wget -nc https://dl.winehq.org/wine-builds/winehq.key \
    && apt-key add winehq.key \
    && apt-add-repository https://dl.winehq.org/wine-builds/debian/ \
    && apt update \
    && wget -O- -q https://download.opensuse.org/repositories/Emulators:/Wine:/Debian/Debian_10/Release.key | apt-key add -  \  
    && echo "deb http://download.opensuse.org/repositories/Emulators:/Wine:/Debian/Debian_10 ./" | tee /etc/apt/sources.list.d/wine-obs.list \
    && DEBIAN_FRONTEND="noninteractive" apt-get install -y --install-recommends winehq-${WINE_BRANCH} \
    && apt-get clean \
    && apt-get autoremove
    
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
#volume for mounting games
VOLUME [ "/games" ]

#
# copy configuation files
COPY supervisord.conf /etc/neko/supervisord/johnwick.conf
# COPY openbox.xml /etc/neko/openbox.xml
