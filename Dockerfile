FROM scratch
COPY assets/certs /etc/ssl/certs
COPY assets/group assets/passwd assets/shadow assets/profile /etc/
COPY /assets/busybox /bin/sh
COPY /assets/busybox /bin/wget
ADD assets/gccbase.tar /
ADD assets/libgcc.tar /
ADD assets/libc6.tar /
# TODO: better location for this?
COPY strato /sbin/
RUN strato add busybox
# TODO: permission on these?
#RUN mkdir -p /bin /sbin /usr/bin /usr/sbin
#RUN touch /etc/sudoers
# TODO: make file layout in one layer
#RUN mkdir /home
