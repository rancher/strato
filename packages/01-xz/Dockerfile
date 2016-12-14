FROM ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
ENV VERSION 5.2.2
RUN wget http://tukaani.org/xz/xz-${VERSION}.tar.gz
RUN tar xf /xz*
RUN cd /xz* \
    && ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --mandir=/usr/share/man \
    --infodir=/usr/share/info \
    --localstatedir=/var \
    --disable-rpath \
    --disable-werror \
    && make
RUN echo '{}' > /_package.yml && cd /xz* \
    && make install
