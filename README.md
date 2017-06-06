# strato

strato is a package manager and minimal container base image. All packages in strato are created by a simple and containerized build process. The latest version is 0.0.2.

```
docker run -it --rm rancher/strato:0.0.2 sh
/ # strato --source="https://github.com/rancher/strato-packages/raw/master/0.0.2/$(get-arch)/" add bash
Installing package https://github.com/rancher/strato-packages/raw/master/0.0.2/amd64/bash.tar.gz:dev/bin
/bin/bash
/bin/bashbug
/usr
/usr/share
1.015647 mb
/ # bash
bash-4.3#
```

## How packages are built

The build instructions for packages are described using a Dockerfile. The package will be built by extracting the files in the final layer from the build process.

As an example, the following Dockerfile builds the GNU make package.

```
FROM ubuntu
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN wget -P /usr/src/ ftp://ftp.gnu.org/gnu/make/make-4.2.1.tar.bz2
RUN cd /usr/src/ && tar xf make*
RUN cd /usr/src/make* \
    && ./configure \
    --prefix=/usr \
    --mandir=/usr/share/man \
    --infodir=/usr/share/info \
    --disable-nls \
    && make

# The following layer is extracted to generate the resulting package
RUN cd /usr/src/make* \
    && make install
```

All packages are currently built using Ubuntu 16.04 as the base image, but the goal is to eventually have enough packages for strato to be able to build itself.

### How to build the packages

The base Dockerfiles are in the packages directory in this repository. To build them, you need `dapper` - from https://github.com/rancher/dapper/releases

Run `dapper build-bin` to build the binaries in this repo, then build the tarballs using `dapper build-packages`
The tarballs will be in the `dist` dir - and can then be uploaded to somewhere that you can get to them.

## Base image

The strato base image includes busybox, glibc, and the strato package manager.
