FROM golang:1.15-alpine

WORKDIR /app

EXPOSE 8000

CMD ["go","run","."]