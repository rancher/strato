FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev liblzma-dev
RUN apt-get update && apt-get install -y autoconf automake libtool
ENV VERSION 0.12.1
ENV LDFLAGS -s
RUN wget -P /usr/src/ https://s3.amazonaws.com/json-c_releases/releases/json-c-${VERSION}.tar.gz
RUN cd /usr/src/ && tar xf json-c*
RUN cd /usr/src/json-c* \
    && ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --mandir=/usr/share/man \
    --infodir=/usr/share/info \
    --localstatedir=/var \
    --disable-static \
    --enable-shared \
    && make

RUN cd /usr/src/json-c* \
    && make install
