FROM golang:1.15-alpine
RUN apk update && apk add git && go get github.com/codegangsta/gin
WORKDIR /app
EXPOSE 8000
CMD ["gin","--appPort","8000","--all","-i","run","main.go"]