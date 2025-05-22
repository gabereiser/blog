FROM golang:alpine AS builder

RUN apk add --no-cache make gcc libc-dev

RUN mkdir /opt/blog

WORKDIR /opt/blog/

COPY ../. /opt/blog/

RUN make all

FROM golang:alpine

WORKDIR /opt
COPY --from=builder /opt/blog/bin/blog /opt/blog
COPY --from=builder /opt/blog/.env /opt/.env

EXPOSE 8080
CMD ["/opt/blog"]