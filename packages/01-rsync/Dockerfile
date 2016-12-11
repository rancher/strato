# TODO: remove dependency on attr
FROM ubuntu:16.10
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y perl acl-dev libpopt-dev
ENV VERSION 3.1.2
RUN wget http://rsync.samba.org/ftp/rsync/rsync-${VERSION}.tar.gz
RUN tar xf /rsync*
RUN cd /rsync* \
    && ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --mandir=/usr/share/man \
    --localstatedir=/var \
    && make
RUN echo '{}' > /_package.yml && cd /rsync* \
    && make install
