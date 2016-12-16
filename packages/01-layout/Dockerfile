FROM ubuntu:16.04
RUN apt-get update && apt-get install -y wget
RUN mkdir assets \
    && cd assets \
    && wget https://raw.githubusercontent.com/alpinelinux/aports/master/main/alpine-baselayout/group \
    && wget https://github.com/alpinelinux/aports/blob/master/main/alpine-baselayout/passwd \
    && wget https://raw.githubusercontent.com/alpinelinux/aports/master/main/alpine-baselayout/profile
COPY shadow assets/

COPY strato.yml /
RUN install -m 0755 -d \
    etc \
    run \
    sbin \
    usr/bin \
    usr/local/bin \
    usr/local/lib \
    usr/local/share \
    usr/sbin \
    usr/share \
    usr/share/man \
    usr/share/misc \
    var/cache \
    var/cache/misc \
    var/lib \
    var/lib/misc \
    var/local \
    var/log \
    var/opt \
    var/run \
    && install -d -m 0555 var/empty \
    && install -d -m 0700 root \
    && install -d -m 1777 tmp var/tmp \
    && install -m644 assets/group etc/group \
    && install -m644 assets/passwd etc/passwd \
    && install -m644 assets/profile etc/profile \
    && install -m644 assets/shadow etc/shadow \
