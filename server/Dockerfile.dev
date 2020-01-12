FROM golang:1.12.7-alpine3.9

WORKDIR /api

ENV GO111MODULE=on

RUN apk add --no-cache \
        alpine-sdk \
        git \
    && go get github.com/pilu/fresh

EXPOSE 8000

CMD ["fresh"]
