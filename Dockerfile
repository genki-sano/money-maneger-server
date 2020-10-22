FROM golang:alpine as builder

RUN apk update \
  && apk add --no-cache git curl make gcc g++

RUN go get -u github.com/cosmtrek/air \
  && chmod +x /go/bin/air

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

ENV API_REVISION=release
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o /main

FROM alpine:3.9

COPY --from=builder /main .

ENV PORT=${PORT}
ENTRYPOINT ["/main web"]
