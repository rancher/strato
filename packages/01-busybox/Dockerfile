FROM ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential wget
RUN wget https://busybox.net/downloads/busybox-1.25.1.tar.bz2
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
