FROM ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev liblzma-dev
ENV VERSION 23
RUN wget http://ftp.kernel.org/pub/linux/utils/kernel/kmod/kmod-${VERSION}.tar.xz
RUN tar xf /kmod*
RUN cd /kmod* \
    && ./configure \
    --prefix=/usr \
    --bindir=/bin \
    --sysconfdir=/etc \
    --with-rootlibdir=/lib \
    --with-zlib \
    --with-xz \
    && make
RUN echo '{}' > /_package.yml && cd /kmod* \
    && make install \
    && for i in lsmod rmmod insmod modinfo modprobe depmod; do [ -e /bin/$i ] && rm /bin/$i; ln -s kmod /bin/$i; done \
    && for i in lsmod modinfo; do [ -e /bin/$i ] && rm /bin/$i; ln -s kmod "$pkgdir"/bin/$i; done
