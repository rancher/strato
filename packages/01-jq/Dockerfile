FROM ubuntu:16.10
RUN apt-get update && apt-get install -y build-essential wget
ENV VERSION 1.5
RUN wget https://github.com/stedolan/jq/releases/download/jq-${VERSION}/jq-${VERSION}.tar.gz
RUN tar xf /jq*
RUN cd /jq* \
    && ./configure --prefix=/usr --disable-docs \
    && make
RUN echo '{}' > /_package.yml && cd /jq* \
    && make prefix=/usr install
