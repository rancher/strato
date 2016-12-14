FROM ubuntu:16.04
# Borrow OpenSSL from Ubuntu for now
COPY libssl.tar /
COPY openssl.tar /

COPY strato.yml /
RUN tar xvf /libssl.tar -C / \
    && tar xvf /openssl.tar -C /
