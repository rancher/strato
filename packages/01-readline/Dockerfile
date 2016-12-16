FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
ENV VERSION 7.0
ENV LDFLAGS -s
RUN wget -P /usr/src/ https://ftp.gnu.org/gnu/readline/readline-${VERSION}.tar.gz
RUN cd /usr/src/ && tar xf readline*
RUN cd /usr/src/readline* \
    && ./configure \
    --prefix=/usr \
    --mandir=/usr/share/man \
    --infodir=/usr/share/info \
    --disable-static \
    --enable-shared \
    && make

COPY strato.yml /
RUN cd /usr/src/readline* \
    && make install
