FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
ENV VERSION 0.1.10
ENV LDFLAGS -s
RUN wget -P /usr/src/ http://libestr.adiscon.com/files/download/libestr-${VERSION}.tar.gz
RUN cd /usr/src/ && tar xf libestr*
RUN cd /usr/src/libestr* \
    && ./configure \
    --prefix=/usr \
    && make

RUN cd /usr/src/libestr* \
    && make install
