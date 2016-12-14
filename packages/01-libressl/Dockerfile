FROM ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential wget
RUN wget https://ftp.openbsd.org/pub/OpenBSD/LibreSSL/libressl-2.5.0.tar.gz
RUN tar xf /libressl*
RUN cd /libressl* \
    && ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --mandir=/usr/share/man \
    --localstatedir=/var \
    && make

COPY strato.yml /
# TODO: remove 1.0.0 first and fix weird permission error
RUN cd /libressl* \
    && make install \
    #&& rm /usr/lib/libcrypto.so.1.0.0 \
    && ln -s /usr/lib/libcrypto.so.38.0.1 /usr/lib/libcrypto.so.1.0.0
