FROM golang:alpine AS builder

RUN mkdir /opt/blog

WORKDIR /opt/blog

COPY . /opt/blog/

RUN make

FROM golang:alpine

RUN mkdir /opt
WORKDIR /opt
COPY --from=builder /opt/blog/bin/blog /opt/blog

EXPOSE 8080
CMD ["/opt/blog"]