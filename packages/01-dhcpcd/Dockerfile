FROM ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
ENV VERSION 6.11.5
ENV LDFLAGS -s
RUN wget -P /usr/src/ http://roy.marples.name/downloads/dhcpcd/dhcpcd-${VERSION}.tar.xz
RUN cd /usr/src/ && tar xf dhcpcd*
RUN cd /usr/src/dhcpcd* \
    && ./configure \
    --sysconfdir=/etc \
    --mandir=/usr/share/man \
    --localstatedir=/var \
    --libexecdir=/usr/lib/dhcpcd \
    --dbdir=/var/lib/dhcpcd \
    && make

COPY strato.yml /
RUN cd /usr/src/dhcpcd* \
    && make install
