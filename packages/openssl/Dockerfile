FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=skip arm=armhf/ubuntu:16.04
# Borrow OpenSSL from Ubuntu for now
COPY libssl.tar /
COPY openssl.tar /

COPY strato.yml /
RUN tar xvf /libssl.tar -C / \
    && tar xvf /openssl.tar -C /
