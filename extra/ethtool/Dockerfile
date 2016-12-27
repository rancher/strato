FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
ENV VERSION 4.8
ENV LDFLAGS -s
RUN wget -P /usr/src/ http://www.kernel.org/pub/software/network/ethtool/ethtool-${VERSION}.tar.xz
RUN cd /usr/src/ && tar xf ethtool*
RUN cd /usr/src/ethtool* \
    && ./configure \
    --prefix=/usr \
    --mandir=/usr/share/man \
    && make

COPY strato.yml /
RUN cd /usr/src/ethtool* \
    && make install
