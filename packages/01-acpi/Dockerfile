FROM ubuntu:16.10
RUN apt-get update && apt-get install -y build-essential wget
ENV VERSION 1.7
RUN wget http://downloads.sourceforge.net/acpiclient/acpi-${VERSION}.tar.gz
RUN tar xf /acpi*
RUN cd /acpi* && ./configure --prefix=/usr
RUN cd /acpi* \
    && make install
