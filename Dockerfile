FROM golang:1.20 as build
LABEL description="Build container"

WORKDIR /build
ENV CGO_ENABLED 0

# Download necessary Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -ldflags "-s" github.com/paralus/relay

FROM alpine:latest as runtime
LABEL description="Run container"

WORKDIR /usr/bin
COPY --from=build /build/relay /usr/bin/relay

ENTRYPOINT ["./relay"]
CMD ["--help"]

# server port
EXPOSE 10002
