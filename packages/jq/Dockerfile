FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential wget
ENV VERSION 1.5
ENV LDFLAGS -s
RUN wget -P /usr/src/ https://github.com/stedolan/jq/releases/download/jq-${VERSION}/jq-${VERSION}.tar.gz
RUN cd /usr/src/ && tar xf jq*
RUN cd /usr/src/jq* \
    && ./configure --prefix=/usr --disable-docs \
    && make

COPY strato.yml /
RUN cd /usr/src/jq* \
    && make prefix=/usr install
