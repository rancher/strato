FROM ubuntu:16.10
RUN apt-get update && apt-get install -y build-essential pkg-config wget
ENV VERSION 6.11.5
RUN wget http://roy.marples.name/downloads/dhcpcd/dhcpcd-${VERSION}.tar.xz
RUN tar xf /dhcpcd*
RUN cd /dhcpcd* \
    && ./configure \
    --sysconfdir=/etc \
    --mandir=/usr/share/man \
    --localstatedir=/var \
    --libexecdir=/usr/lib/dhcpcd \
    --dbdir=/var/lib/dhcpcd \
    && make
RUN echo '{}' > /_package.yml && cd /dhcpcd* \
    && make install
