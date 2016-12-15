# TODO: remove dependency on attr
FROM ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y perl acl-dev libpopt-dev
ENV VERSION 3.1.2
ENV LDFLAGS -s
RUN wget -P /usr/src/ http://rsync.samba.org/ftp/rsync/rsync-${VERSION}.tar.gz
RUN cd /usr/src/ && tar xf rsync*
RUN cd /usr/src/rsync* \
    && ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --mandir=/usr/share/man \
    --localstatedir=/var \
    && make

COPY strato.yml /
RUN cd /usr/src/rsync* \
    && make install
