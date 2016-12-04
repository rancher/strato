FROM alpine
RUN apk --no-cache add gcc make musl-dev
RUN wget http://zlib.net/zlib-1.2.8.tar.gz
RUN tar xf /zlib*
RUN cd /zlib* \
    && ./configure --prefix=/usr --libdir=/lib --shared \
    && make
RUN echo '{}' > /_package.yml && cd /zlib* \
    && make install
