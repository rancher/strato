FROM alpine
RUN apk --no-cache add wget gcc make
RUN wget --no-check-certificate https://www.musl-libc.org/releases/musl-1.1.15.tar.gz
RUN tar xf /musl*
RUN cd /musl* \
    && ./configure --prefix=/usr --sysconfdir=/etc --mandir=/usr/share/man --infodir=/usr/share/info --localstatedir=/var \
    && make
RUN echo '{}' > /_package.yml && cd /musl* \
    && make install
