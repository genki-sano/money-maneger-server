FROM golang:alpine as builder

RUN apk update \
  && apk add --no-cache git curl \
  && go get -u github.com/cosmtrek/air \
  && go get -u github.com/google/wire/cmd/wire \
  && chmod +x ${GOPATH}/bin/air \
  && chmod +x ${GOPATH}/bin/wire

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/package/infrastructure/di
RUN wire

WORKDIR /app
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /main ./package

FROM alpine:3.9

COPY --from=builder /main .

ENV PORT=${PORT}
ENTRYPOINT ["/main web"]
