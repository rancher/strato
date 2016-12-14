FROM ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev liblzma-dev
ENV VERSION 4.2.1
RUN wget ftp://ftp.gnu.org/gnu/make/make-${VERSION}.tar.bz2
RUN tar xf /make*
RUN cd /make* \
    && ./configure \
    --prefix=/usr \
    --mandir=/usr/share/man \
    --infodir=/usr/share/info \
    --disable-nls \
    && make
RUN echo '{}' > /_package.yml && cd /make* \
    && make install
