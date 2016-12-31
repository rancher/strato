FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential wget
ENV LDFLAGS -s
RUN wget -P /usr/src/ http://zlib.net/zlib-1.2.8.tar.gz
RUN cd /usr/src/ && tar xf zlib*
RUN cd /usr/src/zlib* \
    && ./configure \
    --prefix=/usr \
    --libdir=/lib \
    --shared \
    && make

RUN cd /usr/src/zlib* \
    && make install
