FROM golang:latest as build-artifact
WORKDIR /go/src/github.com/mono83/charlie
COPY . .
RUN make build

FROM alpine:latest
RUN apk add ca-certificates
WORKDIR /charlie
COPY --from=build-artifact /go/src/github.com/mono83/charlie/release .
COPY --from=build-artifact /go/src/github.com/mono83/charlie/config.ini .
CMD ["./runner"]