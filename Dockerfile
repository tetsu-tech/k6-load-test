FROM golang:1.19-alpine3.16 as builder
WORKDIR /var/www
COPY . .
RUN go build -o app .

FROM alpine:3.16 as app
WORKDIR /var/www
COPY --from=builder /var/www/app /usr/local/bin/app

EXPOSE 8080
ENTRYPOINT ["app"]
