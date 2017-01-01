FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev liblzma-dev
RUN apt-get update && apt-get install -y libglib2.0-dev dh-autoreconf libmspack-dev libdumbnet-dev libssl-dev
ENV VERSION 10.0.7
ENV LDFLAGS -s
RUN wget -P /usr/src/ https://github.com/vmware/open-vm-tools/archive/stable-${VERSION}.tar.gz
RUN cd /usr/src/ && tar xf stable*

RUN cd /usr/src/open-vm-tools*/open-vm-tools \
    && autoreconf -vif \
    && ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --without-xerces \
    --without-icu \
    --without-procps \
    --without-x \
    --without-kernel-modules \
    --without-pam \
    && make

RUN cd /usr/src/open-vm-tools*/open-vm-tools \
    && make install
