FROM ubuntu:16.10
RUN apt-get update && apt-get install -y build-essential pkg-config wget
ENV VERSION 1.4.17
RUN wget ftp://ftp.gnu.org/gnu/m4/m4-${VERSION}.tar.gz
RUN tar xf /m4*
RUN cd /m4* \
    && ./configure \
    --prefix=/usr \
    && make
RUN echo '{}' > /_package.yml && cd /m4* \
    && make install
