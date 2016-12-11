FROM ubuntu:16.10
# Borrow OpenSSL from Ubuntu for now
COPY libssl.tar /
COPY openssl.tar /
RUN echo '{}' > /_package.yml \
    && tar xvf /libssl.tar -C / \
    && tar xvf /openssl.tar -C /
