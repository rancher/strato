FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
ENV VERSION 4.6.0
ENV LDFLAGS -s
RUN wget -P /usr/src/ http://ftp.gnu.org/pub/gnu/findutils/findutils-${VERSION}.tar.gz
RUN cd /usr/src/ && tar xf findutils*
RUN cd /usr/src/findutils* \
    && ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --localstatedir=/var \
    --mandir=/usr/share/man \
    --infodir=/usr/share/info \
    && make

COPY strato.yml /
RUN cd /usr/src/findutils* \
    && make install
