FROM alpine:3.8

EXPOSE 3000

RUN mkdir /app \
&&  mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 \
&&  apk update && apk add ca-certificates

COPY build/weather-server /app

WORKDIR /app

ENTRYPOINT ["./weather-server"]
