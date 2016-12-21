FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev libssl-dev
ENV LDFLAGS -s
RUN wget -P /usr/src/ http://www.eecis.udel.edu/~ntp/ntp_spool/ntp4/ntp-4.2/ntp-4.2.8p9.tar.gz
COPY ntp.conf config.guess config.sub /usr/src/
RUN cd /usr/src/ \
    && tar xf ntp*.tar.gz \
    && cd ntp* \
    && cp ../config.guess ../config.sub sntp/libevent/build-aux
RUN cd /usr/src/ntp* \
    && ./configure \
    --prefix=/usr \
    --bindir=/usr/sbin \
    && make

COPY strato.yml /
RUN cd /usr/src/ntp* \
    && make install \
    && install -m644 -D /usr/src/ntp.conf /etc/ntp.conf
