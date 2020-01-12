FROM golang:1.12.7 as builder

WORKDIR /go/src/app
ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

FROM alpine
COPY --from=builder /go/bin/app /go/bin/app

CMD /go/bin/app $PORT
