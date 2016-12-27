FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev liblzma-dev
RUN apt-get update && apt-get install -y libglib2.0-dev
ENV VERSION 7.3.0
ENV LDFLAGS -s
RUN wget -P /usr/src/ https://nodejs.org/dist/v${VERSION}/node-v${VERSION}.tar.xz
RUN cd /usr/src/ && tar xf node*
RUN cd /usr/src/node* \
    && ./configure \
    --prefix=/usr \
    && make

RUN mv /usr/lib/x86_64-linux-gnu/libstdc++* /

COPY strato.yml /
RUN cd /usr/src/node* \
    && make install \
    && cp /libstdc++* /usr/lib/x86_64-linux-gnu/
