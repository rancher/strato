FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y bison flex iptables-dev libnftnl-dev libmnl-dev
ENV VERSION 1.6.0
ENV LDFLAGS -s
RUN wget -P /usr/src/ http://ftp.netfilter.org/pub/iptables/iptables-${VERSION}.tar.bz2
RUN cd /usr/src/ && tar xf iptables*
RUN cd /usr/src/iptables* \
    && ./configure \
    --prefix=/usr \
    --mandir=/usr/share/man \
    --sbindir=/sbin \
    --without-kernel \
    --enable-devel \
    --enable-libipq \
    --enable-shared \
    && make

RUN cd /usr/src/iptables* \
    && make install
