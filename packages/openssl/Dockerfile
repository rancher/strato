FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
# Borrow OpenSSL from Ubuntu for now
COPY assets /

COPY strato.yml /
RUN tar xvf "$(uname -m)/libssl.tar" -C / \
    && tar xvf "$(uname -m)/openssl.tar" -C /
