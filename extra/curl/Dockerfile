FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
ENV VERSION 7.51.0
ENV LDFLAGS -s
RUN wget -P /usr/src/ http://curl.haxx.se/download/curl-${VERSION}.tar.bz2
RUN cd /usr/src/ && tar xf curl*
RUN cd /usr/src/curl* \
    && ./configure \
    --prefix=/usr \
    --enable-ipv6 \
    --enable-unix-sockets \
    --without-libidn \
    --without-libidn2 \
    --disable-static \
    --disable-ldap \
    --with-pic \
    && make

RUN cd /usr/src/curl* \
    && make install
