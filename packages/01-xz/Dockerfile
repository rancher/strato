FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
ENV VERSION 5.2.2
ENV LDFLAGS -s
RUN wget -P /usr/src/ http://tukaani.org/xz/xz-${VERSION}.tar.gz
RUN cd /usr/src/ && tar xf xz*
RUN cd /usr/src/xz* \
    && ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --mandir=/usr/share/man \
    --infodir=/usr/share/info \
    --localstatedir=/var \
    --disable-rpath \
    --disable-werror \
    && make

COPY strato.yml /
RUN cd /usr/src/xz* \
    && make install
