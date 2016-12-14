FROM ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev libssl-dev
RUN wget http://www.eecis.udel.edu/~ntp/ntp_spool/ntp4/ntp-4.2/ntp-4.2.8p9.tar.gz
RUN tar xf /ntp*
RUN cd /ntp* \
    && ./configure \
    --prefix=/usr \
    --bindir=/usr/sbin \
    && make

COPY strato.yml /
RUN cd /ntp* \
    && make install
