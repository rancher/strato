FROM alpine
RUN apk --no-cache add gcc make musl-dev linux-headers curl
RUN curl https://ftp.openbsd.org/pub/OpenBSD/LibreSSL/libressl-2.5.0.tar.gz > /libressl.tar.gz
RUN tar xf /libressl*
RUN cd /libressl* \
    && ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --mandir=/usr/share/man \
    --localstatedir=/var \
    && make
# TODO: remove 1.0.0 first and fix weird permission error
RUN echo '{}' > /_package.yml && cd /libressl* \
    && make install \
    && rm /usr/lib/libcrypto.so.1.0.0 \
    && ln -s /usr/lib/libcrypto.so.38.0.1 /usr/lib/libcrypto.so.1.0.0
