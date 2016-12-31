FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev liblzma-dev
RUN apt-get update && apt-get install -y libestr-dev libjson-c-dev uuid-dev
ENV VERSION 8.16.0
ENV LDFLAGS -s
RUN wget -P /usr/src/ http://www.rsyslog.com/files/download/rsyslog/rsyslog-${VERSION}.tar.gz
COPY rsyslog.conf /usr/src/
RUN cd /usr/src/ && tar xf rsyslog*.tar.gz
RUN cd /usr/src/rsyslog* \
    && ./configure \
    --prefix=/usr \
    --disable-rfc3195 \
    --enable-largefile \
    --disable-imdiag \
    --disable-imfile \
    --disable-mail \
    --disable-omprog \
    --disable-omstdout \
    --disable-omudpspoof \
    --disable-imptcp \
    --disable-impstats \
    --disable-mysql \
    --disable-pgsql \
    --disable-gnutls \
    --disable-snmp \
    --disable-omhiredis \
    --disable-libgcrypt \
    --disable-liblogging-stdlog \
    && make

RUN cd /usr/src/rsyslog* \
    && make install \
    && install -m644 -D /usr/src/rsyslog.conf /etc/rsyslog.conf
