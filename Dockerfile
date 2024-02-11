FROM golang:1.22.0-alpine3.18

RUN apk update && apk add git vim

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

WORKDIR /backend

COPY /backend/* ./

RUN go mod download

RUN go install github.com/cosmtrek/air@latest

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
