FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev liblzma-dev
ENV VERSION 4.2.1
ENV LDFLAGS -s
RUN wget -P /usr/src/ ftp://ftp.gnu.org/gnu/make/make-${VERSION}.tar.bz2
RUN cd /usr/src/ && tar xf make*
RUN cd /usr/src/make* \
    && ./configure \
    --prefix=/usr \
    --mandir=/usr/share/man \
    --infodir=/usr/share/info \
    --disable-nls \
    && make

RUN cd /usr/src/make* \
    && make install
