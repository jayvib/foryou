# Build
FROM golang:1.12.8-alpine3.10 as builder

RUN apk update && apk upgrade && \
    apk --update add git gcc make


WORKDIR /go/src/github.com/jayvib/app

COPY . .

ENV GO111MODULE on

RUN go build -o engine.linux

# Distribute
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    apk add --no-cache bash && \
    mkdir /app

WORKDIR /app

EXPOSE 8080

COPY --from=builder /go/src/github.com/jayvib/app/engine.linux /app

CMD /app/engine.linux
