FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
ENV VERSION 6.0
RUN wget https://ftp.gnu.org/gnu/ncurses/ncurses-${VERSION}.tar.gz
RUN tar xf /ncurses*
RUN cd /ncurses* \
    && CPPFLAGS=-P ./configure \
    --with-shared   \
    --without-debug \
    --without-ada   \
    --enable-widec  \
    --enable-overwrite \
    && make

COPY strato.yml /
RUN cd /ncurses* \
    && make install
