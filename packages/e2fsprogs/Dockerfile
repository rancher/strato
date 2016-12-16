FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y uuid-dev libblkid-dev
ENV VERSION 1.43.3
RUN wget -P /usr/src/ https://www.kernel.org/pub/linux/kernel/people/tytso/e2fsprogs/v${VERSION}/e2fsprogs-${VERSION}.tar.xz
RUN cd /usr/src/ && tar xf e2fsprogs*
ENV LDFLAGS -s
RUN cd /usr/src/e2fsprogs* \
    && ./configure \
    --mandir=/usr/share/man \
    --enable-elf-shlibs \
    --enable-symlink-install \
    --disable-fsck \
    --disable-uuidd \
    --disable-libuuid \
    --disable-libblkid \
    --disable-tls \
    --disable-nls \
    && make

COPY strato.yml /
RUN cd /usr/src/e2fsprogs* \
    && make install install-libs \
