FROM alpine:latest as builder

LABEL maintainer="Danil Korennykh <bluskript@gmail.com>"

RUN apk add git go vips vips-dev

WORKDIR /legato

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o legato .

EXPOSE 2289

CMD [ "./legato" ]