FROM ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev liblzma-dev
ENV VERSION 2.29
RUN wget https://www.kernel.org/pub/linux/utils/util-linux/v${VERSION}/util-linux-${VERSION}.tar.gz
RUN tar xf /util-linux*
RUN cd /util-linux* \
    && ./configure \
    --prefix=/usr \
    --without-python               \
    --without-systemdsystemunitdir \
    && make

COPY strato.yml /
RUN cd /util-linux* \
    && make install
