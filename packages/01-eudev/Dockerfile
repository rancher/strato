FROM ubuntu:16.10
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y gperf libglib2.0-dev libkmod-dev
ENV VERSION 3.2
RUN wget http://dev.gentoo.org/~blueness/eudev/eudev-${VERSION}.tar.gz
RUN tar xf /eudev*
RUN cd /eudev* \
    && ./configure \
    --sysconfdir=/etc \
    --with-rootprefix= \
    --with-rootrundir=/run \
    --with-rootlibexecdir=/lib/udev \
    --libdir=/usr/lib \
    --enable-split-usr \
    --enable-manpages \
    --disable-hwdb \
    --enable-kmod \
    --exec-prefix=/ \
    && make
RUN echo '{}' > /_package.yml && cd /eudev* \
    && make sharepkgconfigdir=/usr/lib/pkgconfig install
