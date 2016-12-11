FROM ubuntu:16.10
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y bison flex iptables-dev
ENV VERSION 4.7.0
RUN wget http://kernel.org/pub/linux/utils/net/iproute2/iproute2-${VERSION}.tar.xz
RUN tar xf /iproute2*
RUN cd /iproute2* \
    && ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --mandir=/usr/share/man \
    --localstatedir=/var \
    && make
RUN echo '{}' > /_package.yml && cd /iproute2* \
    && make install
