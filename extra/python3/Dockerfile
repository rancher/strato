FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev liblzma-dev
RUN apt-get update && apt-get install -y libglib2.0-dev
ENV VERSION 3.5.2
ENV LDFLAGS -s
RUN wget -P /usr/src/ http://www.python.org/ftp/python/${VERSION}/Python-${VERSION}.tar.xz
RUN cd /usr/src/ && tar xf Python*
RUN cd /usr/src/Python* \
    && ./configure \
    --prefix=/usr \
    --disable-rpath \
    --enable-ipv6 \
    --enable-loadable-sqlite-extensions \
    --enable-shared \
    --with-computed-gotos \
    --with-dbmliborder=gdbm:ndbm \
    --with-system-expat \
    --with-system-ffi \
    --with-threads \
    && make

RUN cd /usr/src/Python* \
    && make install
