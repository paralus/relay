FROM golang:1.17 as build
LABEL description="Build container"

ENV CGO_ENABLED 0
COPY . /build
WORKDIR /build
RUN go build github.com/RafaySystems/relay

FROM alpine:latest as runtime
LABEL description="Run container"

COPY --from=build /build/relay /usr/bin/relay
WORKDIR /usr/bin
ENTRYPOINT ["./relay"]
CMD ["--help"]

EXPOSE 13000
