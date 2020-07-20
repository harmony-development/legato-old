# adapted from https://levelup.gitconnected.com/dockerized-crud-restful-api-with-go-gorm-jwt-postgresql-mysql-and-testing-61d731430bd8

FROM golang:alpine as builder

LABEL maintainer="Danil Korennykh <bluskript@gmail.com>"

RUN apk update && apk add --no-cache git
RUN apk --no-cache add vips-dev fftw-dev build-base

WORKDIR /legato

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux go build -a -installsuffix cgo -o legato .
