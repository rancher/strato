FROM ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y uuid-dev libblkid-dev
ENV VERSION 1.43.3
RUN wget https://www.kernel.org/pub/linux/kernel/people/tytso/e2fsprogs/v${VERSION}/e2fsprogs-${VERSION}.tar.xz
RUN tar xf /e2fsprogs*
RUN cd /e2fsprogs* \
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
RUN echo '\n\
exclude:\n\
- usr/share/man\n'\
>> /_package.yml \
    && cd /e2fsprogs* \
    && make install install-libs \
