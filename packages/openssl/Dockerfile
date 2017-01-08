FROM strato-build

# Borrow OpenSSL from Ubuntu for now
COPY assets /

RUN tar xvf "$(uname -m)/libssl.tar" -C / \
    && tar xvf "$(uname -m)/openssl.tar" -C /
