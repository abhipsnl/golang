FROM golang:1.15-alpine as dev

RUN apk --update add --no-cache ca-certificates openssl git tzdata && \
update-ca-certificates

WORKDIR /work

FROM golang:1.15-alpine as build

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
WORKDIR /redis
COPY ./redis/* /redis/
RUN go build -o redis

FROM alpine as runtime 

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /redis/redis /
CMD ./redis
