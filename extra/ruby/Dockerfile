FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev liblzma-dev libssl-dev
ENV VERSION 2.3.3
ENV LDFLAGS -s
RUN wget -P /usr/src/ https://cache.ruby-lang.org/pub/ruby/2.3/ruby-${VERSION}.tar.gz
RUN cd /usr/src/ && tar xf ruby*
RUN cd /usr/src/ruby* \
    && ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --mandir=/usr/share/man \
    --infodir=/usr/share/info \
    --with-sitedir='/usr/local/lib/site_ruby' \
    --enable-pthread \
    --disable-rpath \
    --enable-shared \
    --with-mantype=man \
    && make

COPY strato.yml /
RUN cd /usr/src/ruby* \
    && make install
