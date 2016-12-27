FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev liblzma-dev
RUN apt-get update && apt-get install -y libglib2.0-dev autoconf
ENV VERSION 1.5
ENV LDFLAGS -s
RUN wget -P /usr/src/ http://downloads.sourceforge.net/project/bridge/bridge/bridge-utils-${VERSION}.tar.gz
RUN cd /usr/src/ && tar xf bridge*
RUN cd /usr/src/bridge* \
    && autoconf \
    && ./configure \
    --prefix=/usr \
    --mandir=/usr/share/man \
    --libdir=/usr/lib \
    --includedir=/usr/include \
    && make

COPY strato.yml /
RUN cd /usr/src/bridge* \
    && make install
