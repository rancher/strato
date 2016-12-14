FROM ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev liblzma-dev
ENV VERSION 4.4
RUN wget https://ftp.gnu.org/gnu/bash/bash-${VERSION}.tar.gz
RUN tar xf /bash*
RUN cd /bash* \
    && ./configure \
    --prefix=/usr \
    --bindir=/bin \
    --without-bash-malloc \
    && make
RUN echo '{}' > /_package.yml && cd /bash* \
    && make install
