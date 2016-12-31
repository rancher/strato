FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev liblzma-dev
ENV VERSION 23
ENV LDFLAGS -s
RUN wget -P /usr/src/ http://ftp.kernel.org/pub/linux/utils/kernel/kmod/kmod-${VERSION}.tar.xz
RUN cd /usr/src/ && tar xf kmod*
RUN cd /usr/src/kmod* \
    && ./configure \
    --prefix=/usr \
    --bindir=/bin \
    --sysconfdir=/etc \
    --with-rootlibdir=/lib \
    --with-zlib \
    --with-xz \
    && make

RUN cd /usr/src/kmod* \
    && make install \
    && for i in lsmod rmmod insmod modinfo modprobe depmod; do [ -e /bin/$i ] && rm /bin/$i; ln -s kmod /bin/$i; done \
    && for i in lsmod modinfo; do [ -e /bin/$i ] && rm /bin/$i; ln -s kmod "$pkgdir"/bin/$i; done
