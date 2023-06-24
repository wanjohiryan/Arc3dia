#All we do is copy the ffmpeg from the base docker image, without any of the created users and groups
ARG UBUNTU_RELEASE=22.04
FROM ghcr.io/linuxserver/ffmpeg:amd64-latest as base

FROM ubuntu:${UBUNTU_RELEASE}

ARG DEBIAN_FRONTEND=noninteractive

#
#Install ffmpeg5
RUN mkdir -p etc/OpenCL/vendors
COPY  --from=base /usr/local /usr/local/
COPY  --from=base /etc/OpenCL/vendors/nvidia.icd /etc/OpenCL/vendors/

RUN apt-get update; \
    apt-get install -y \
    libexpat1 libglib2.0-0 libgomp1 libharfbuzz0b libv4l-0 \
    libwayland-client0 libx11-6 libx11-xcb1 libxcb1 libxcb-dri3-0 \
    libxcb-shape0 libxcb-xfixes0 libxext6 libxfixes3 libxml2 ocl-icd-libopencl1; \
    #
    #clean up
    rm -rf /var/lib/apt/lists/* /var/tmp/*

#check the installed version
RUN ffmpeg -version;