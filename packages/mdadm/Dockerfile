FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
ENV VERSION 3.4
ENV LDFLAGS -s
RUN wget -P /usr/src/ http://www.kernel.org/pub/linux/utils/raid/mdadm/mdadm-${VERSION}.tar.gz
RUN cd /usr/src/ && tar xf mdadm*
RUN cd /usr/src/mdadm* \
    && make

COPY strato.yml /
RUN cd /usr/src/mdadm* \
    && make install
