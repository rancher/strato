FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev liblzma-dev
RUN apt-get update && apt-get install -y libglib2.0-dev
ENV VERSION 5.24.0
ENV LDFLAGS -s
RUN wget -P /usr/src/ http://www.cpan.org/src/5.0/perl-${VERSION}.tar.gz
RUN cd /usr/src/ && tar xf perl*
RUN cd /usr/src/perl* \
    && ./Configure -des -Dprefix=/usr \
    && make

COPY strato.yml /
RUN cd /usr/src/perl* \
    && make install
