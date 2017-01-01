FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev liblzma-dev uuid-dev liblvm2-dev libpopt-dev libssl-dev
ENV VERSION 1.7.3
ENV LDFLAGS -s
RUN wget -P /usr/src/ https://www.kernel.org/pub/linux/utils/cryptsetup/v1.7/cryptsetup-${VERSION}.tar.gz
RUN cd /usr/src/ && tar xf cryptsetup*
RUN cd /usr/src/cryptsetup* \
    && ./configure \
    --prefix=/usr \
    --libdir=/lib \
    --sbindir=/sbin \
    --disable-static \
    --with-crypto_backend=openssl \
    && make

RUN cd /usr/src/cryptsetup* \
    && make install
