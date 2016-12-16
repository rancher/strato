FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev liblzma-dev
RUN apt-get update && apt-get install -y uuid-dev libblkid-dev libncurses-dev libreadline-dev liblvm2-dev
ENV VERSION 3.2
ENV LDFLAGS -s
RUN wget -P /usr/src/ ftp://ftp.gnu.org/pub/gnu/parted/parted-${VERSION}.tar.xz
RUN cd /usr/src && tar xf parted*
RUN cd /usr/src/parted* \
    && ./configure \
    --prefix=/usr \
    --disable-debug \
    --disable-nls \
    #--disable-device-mapper \
    --disable-static \
    --enable-shared \
    && make

COPY strato.yml /
RUN cd /usr/src/parted* \
    && make install
