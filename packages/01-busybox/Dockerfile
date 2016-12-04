FROM alpine
RUN apk --no-cache add gcc make curl linux-headers musl-dev
RUN curl https://busybox.net/downloads/busybox-1.25.1.tar.bz2 > busybox.tar.bz2
COPY .config /
RUN tar xf busybox* \
    && cd busybox* \
    && mv /.config .config \
    && make silentoldconfig \
    && make
RUN echo '{}' > /_package.yml \
    && cd /busybox* \
    && make install \
    && install -m 755 busybox /bin/busybox \
    && cd / \
    && for i in $(busybox --list-full); do [ -e $i ] && busybox rm $i; busybox ln -s /bin/busybox $i; done
