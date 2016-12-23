FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev liblzma-dev
RUN apt-get update && apt-get install -y libglib2.0-dev
ENV VERSION 3.14
ENV LDFLAGS -s
RUN wget -P /usr/src/ http://downloads.sourceforge.net/nbd/nbd-${VERSION}.tar.xz
RUN cd /usr/src/ && tar xf nbd*
RUN cd /usr/src/nbd* \
    && touch "systemd/nbd@.service.tmpl" \
    && ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --mandir=/usr/share/man \
    --localstatedir=/var \
    --enable-lfs \
    --enable-syslog \
    --enable-gznbd \
    && make

COPY strato.yml /
RUN cd /usr/src/nbd* \
    && make install
