FROM golang:1.23.1-alpine3.20

WORKDIR /go/src/app/

COPY ./back .

RUN apk upgrade --update && \
    apk --no-cache add git gcc musl-dev