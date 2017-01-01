FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
ENV VERSION 4.89
ENV LDFLAGS -s
RUN wget -P /usr/src/ ftp://lsof.itap.purdue.edu/pub/tools/unix/lsof/lsof_${VERSION}.tar.gz
RUN cd /usr/src/ && tar xf lsof*
RUN cd /usr/src/lsof* \
    && tar xf *src.tar \
    && cd *src \
    && ./Configure -n linux \
    && make

RUN cd /usr/src/lsof*/*src \
    && install -D -m0755 -o root -g root lsof /usr/sbin/lsof \
    && install -D -m0644 -o root -g root lsof.8 /usr/share/man/man8/lsof.8
