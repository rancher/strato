FROM strato-build

# Borrow OpenSSL from Ubuntu for now
COPY assets /

# TODO: bug is causing wrong layer to be extracted if COPY is second to last
RUN true

RUN tar xvf "$(uname -m)/libssl.tar" -C / \
    && tar xvf "$(uname -m)/openssl.tar" -C /
