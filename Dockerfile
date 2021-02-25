# Builder
FROM golang:1.15-alpine as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY . .

RUN make engine

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata mariadb-dev mariadb-client && \
    mkdir /app 

WORKDIR /app 

EXPOSE 8080

COPY --from=builder /app/engine /app
COPY --from=builder /app/scripts/dockerstart.sh /app

CMD /app/dockerstart.sh