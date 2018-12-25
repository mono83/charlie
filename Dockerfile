FROM alpine:latest
RUN apk add ca-certificates
RUN mkdir /charlie
ADD ./release /charlie
ADD ./config.ini /charlie
WORKDIR /charlie
CMD ["./runner"]