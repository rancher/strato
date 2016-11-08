FROM alpine
RUN apk add -U ca-certificates
COPY lay /usr/bin/
