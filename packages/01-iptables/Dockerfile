FROM ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y bison flex iptables-dev libnftnl-dev libmnl-dev
ENV VERSION 1.6.0
RUN wget http://ftp.netfilter.org/pub/iptables/iptables-${VERSION}.tar.bz2
RUN tar xf /iptables*
RUN cd /iptables* \
    && ./configure \
    --prefix=/usr \
    --mandir=/usr/share/man \
    --sbindir=/sbin \
    --without-kernel \
    --enable-devel \
    --enable-libipq \
    --enable-shared \
    && make

COPY strato.yml /
RUN cd /iptables* \
    && make install
