FROM scratch
COPY certs /etc/ssl/certs
COPY group passwd shadow profile /etc/
COPY busybox /bin/sh
# TODO: better location for this?
COPY strato /sbin/
#RUN strato add glibc busybox zlib libressl
ADD glibc.tar /
RUN strato add busybox zlib libressl
# TODO: permission on these?
#RUN mkdir -p /bin /sbin /usr/bin /usr/sbin
#RUN touch /etc/sudoers
# TODO: make file layout in one layer
#RUN mkdir /home
COPY busybox /bin/sh
