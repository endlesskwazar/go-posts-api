FROM golang:1.15-alpine
RUN apk update && apk add git build-base gcc && go get github.com/codegangsta/gin
WORKDIR /go/src/go-cource-api
EXPOSE 8000