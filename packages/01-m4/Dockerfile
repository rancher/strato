FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
ENV VERSION 1.4.17
ENV LDFLAGS -s
RUN wget ftp://ftp.gnu.org/gnu/m4/m4-${VERSION}.tar.gz
RUN tar xf /m4*
RUN cd /m4* \
    && ./configure \
    --prefix=/usr \
    && make

COPY strato.yml /
RUN cd /m4* \
    && make install
