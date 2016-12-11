FROM ubuntu:16.10
RUN apt-get update && apt-get install -y build-essential wget
RUN wget http://zlib.net/zlib-1.2.8.tar.gz
RUN tar xf /zlib*
RUN cd /zlib* \
    && ./configure --prefix=/usr --libdir=/lib --shared \
    && make
RUN echo '{}' > /_package.yml && cd /zlib* \
    && make install
