FROM ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y nasm uuid-dev libblkid-dev
ENV LDFLAGS -s
RUN wget https://www.kernel.org/pub/linux/utils/boot/syslinux/6.xx/syslinux-6.03.tar.xz
RUN tar xf /syslinux*
RUN cd /syslinux* \
    && make

COPY strato.yml /
RUN cd /syslinux* \
    && make install
