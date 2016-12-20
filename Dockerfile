FROM scratch
ADD rootfs.tar /
COPY bin/strato /sbin/
