FROM ubuntu
# FROM amd64=ubuntu:16.04 arm64=aarch64/ubuntu:16.04 arm=armhf/ubuntu:16.04
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN apt-get update && apt-get install -y libssl-dev
ENV VERSION 2.0.873
ENV LDFLAGS -s
RUN wget -P /usr/src/ https://github.com/open-iscsi/open-iscsi/archive/${VERSION}.tar.gz
RUN cd /usr/src/ && tar xf ${VERSION}*
RUN cd /usr/src/open-iscsi* \
    && make user

RUN cd /usr/src/open-iscsi* \
    && make install_user
