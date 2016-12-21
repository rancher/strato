# strato

strato is a package manager and minimal container base image. All packages in strato are created by a containerized build process and then distributed via a container registry.

## How a package is built

The build instructions for packages are described using a Dockerfile and then distributed via Docker Hub. Only the final layer, which actually contains the set of files making up the package, will be installed via strato.

As an example, the following Dockerfile builds the GNU make package.

```
FROM ubuntu
RUN apt-get update && apt-get install -y build-essential pkg-config wget
RUN wget ftp://ftp.gnu.org/gnu/make/make-4.2.1.tar.bz2
RUN tar xf /make*
RUN cd /make* \
    && ./configure \
    --prefix=/usr \
    --mandir=/usr/share/man \
    --infodir=/usr/share/info \
    --disable-nls \
    && make

COPY strato.yml /

# The following container image layer contains all files in the package
RUN cd /make* \
    && make install
```

After this image is built (`docker build -t user/make .`) and pushed to Docker Hub (`docker push user/make`) it can be installed via strato (`strato add user/make`).
