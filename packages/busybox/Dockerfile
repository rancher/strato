FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential wget
ENV LDFLAGS -s
RUN wget -P /usr/src/ https://busybox.net/downloads/busybox-1.25.1.tar.bz2
COPY .config /usr/src/
RUN cd /usr/src/ && tar xf busybox*
RUN cd /usr/src/busybox* \
    && mv ../.config ./ \
    && make silentoldconfig \
    && make

RUN cd /usr/src/busybox* \
    && install -m 755 busybox /bin/busybox \
    && cd / \
    && for i in $(busybox --list-full); do [ -e $i ] && busybox rm $i; busybox ln -s /bin/busybox $i; done
