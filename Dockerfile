FROM golang:alpine as builder

RUN apk update \
  && apk add --no-cache git curl make gcc g++

RUN go get -u github.com/cosmtrek/air \
  && go get -u github.com/google/wire/cmd/wire \
  && chmod +x /go/bin/air \
  && chmod +x /go/bin/wire

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

WORKDIR /app/package/infrastructure/di
RUN wire

WORKDIR /app
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /main ./package

FROM alpine:3.9

COPY --from=builder /main .

ENV PORT=${PORT}
ENTRYPOINT ["/main web"]
