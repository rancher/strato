FROM alpine
RUN apk add -U ca-certificates
COPY strato /usr/bin/
