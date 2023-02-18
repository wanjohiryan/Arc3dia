#! Some code was copied or inspired by https://github.com/selkies-project/docker-nvidia-glx-desktop
# Which is licensed under  MPL-2.0 license
# Thanks ehfd :)

#
#Import from warp, no gpu hw acceleration support yet :(
# Latest stable major version is 0
# Do not use this, it will probably break your game
FROM ghcr.io/wanjohiryan/warp/x264:0

#
#Probably unsafe but we will use the root user from here
#
USER root

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

# ARG VIRTUALGL_VERSION=3.0.2

# #!Needs nvidia cuda base image
# # Install VirtualGL and make libraries available for preload
# RUN curl -fsSL -O "https://sourceforge.net/projects/virtualgl/files/virtualgl_${VIRTUALGL_VERSION}_amd64.deb" && \
#     curl -fsSL -O "https://sourceforge.net/projects/virtualgl/files/virtualgl32_${VIRTUALGL_VERSION}_amd64.deb" && \
#     apt-get update && apt-get install -y --no-install-recommends ./virtualgl_${VIRTUALGL_VERSION}_amd64.deb ./virtualgl32_${VIRTUALGL_VERSION}_amd64.deb && \
#     rm -f "virtualgl_${VIRTUALGL_VERSION}_amd64.deb" "virtualgl32_${VIRTUALGL_VERSION}_amd64.deb" && \
#     rm -rf /var/lib/apt/lists/* && \
#     chmod u+s /usr/lib/libvglfaker.so && \
#     chmod u+s /usr/lib/libdlfaker.so && \
#     chmod u+s /usr/lib32/libvglfaker.so && \
#     chmod u+s /usr/lib32/libdlfaker.so && \
#     chmod u+s /usr/lib/i386-linux-gnu/libvglfaker.so && \
#     chmod u+s /usr/lib/i386-linux-gnu/libdlfaker.so

#
# Install wine >= 7.2 for vk3d capabilities // probably not that stable :)
ARG WINE_BRANCH="staging"
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
#install dx11 and dx9
#! Use dxvk instead, winetricks adds >600mb  
# RUN winetricks -q d3dx9  d3dx11

#
# Silence all fixme warnings from wine
ENV WINEDEBUG fixme-all

#
#volume for mounting games
VOLUME [ "/games" ]

#
#Copy config files
COPY gamevm/base/qwantify /etc/warp/entrypoint.d/brave.sh
RUN chmod +x /etc/warp/entrypoint.d/brave.sh

#
#return back to our designated user
USER ${USERNAME}
