FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential wget
ENV VERSION 1.7
ENV LDFLAGS -s
RUN wget -P /usr/src/ http://downloads.sourceforge.net/acpiclient/acpi-${VERSION}.tar.gz
RUN cd /usr/src/ && tar xf acpi*
RUN cd /usr/src/acpi* \
    && ./configure --prefix=/usr \
    && make

COPY strato.yml /
RUN cd /usr/src/acpi* \
    && make install
