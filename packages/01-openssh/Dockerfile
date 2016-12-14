FROM ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y zlib1g-dev libssl-dev
ENV VERSION 7.3p1
RUN wget http://ftp.openbsd.org/pub/OpenBSD/OpenSSH/portable/openssh-${VERSION}.tar.gz
RUN tar xf /openssh*
RUN cd /openssh* \
    && ./configure \
    --prefix=/usr \
    --sysconfdir=/etc/ssh \
    --datadir=/usr/share/openssh \
    --libexecdir=/usr/lib/ssh \
    --mandir=/usr/share/man \
    --with-pid-dir=/run \
    --with-mantype=man \
    --disable-strip \
    --disable-lastlog \
    --disable-wtmp \
    --with-privsep-path=/var/empty \
    --with-privsep-user=sshd \
    --with-md5-passwords \
    --with-ssl-engine \
    --without-pam \
    && make
RUN echo '{}' > /_package.yml && cd /openssh* \
    && make install
