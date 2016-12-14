FROM ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
ENV VERSION 1.8.18p1
RUN wget https://www.sudo.ws/dist/sudo-${VERSION}.tar.gz
RUN tar xf /sudo*
RUN cd /sudo* \
    && ./configure \
    --prefix=/usr \
    --libexecdir=/usr/lib \
    --mandir=/usr/share/man \
    --disable-nls \
    --enable-pie \
    --without-env-editor \
    --without-pam \
    --without-skey \
    --with-passprompt="[sudo] password for %p: " \
    && make
RUN echo '{}' > /_package.yml && cd /sudo* \
    && make install
