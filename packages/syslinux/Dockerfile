FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=skip arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y nasm uuid-dev libblkid-dev
ENV LDFLAGS -s
RUN wget -P /usr/src/ https://www.kernel.org/pub/linux/utils/boot/syslinux/6.xx/syslinux-6.03.tar.xz
RUN cd /usr/src/ && tar xf syslinux*
RUN cd /usr/src/syslinux* \
    && make

RUN cd /usr/src/syslinux* \
    && make install
