# adapted from https://levelup.gitconnected.com/dockerized-crud-restful-api-with-go-gorm-jwt-postgresql-mysql-and-testing-61d731430bd8

FROM golang:alpine as builder

LABEL maintainer="Danil Korennykh <bluskript@gmail.com>"

RUN apk update && apk add --no-cache git

WORKDIR /harmonyserver

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o harmony-server .

FROM node:12.2.0-alpine as harmonyclientbuilder
WORKDIR /app
RUN apk update && apk add --no-cache git
RUN git clone https://github.com/harmony-development/harmony-app
WORKDIR /app/harmony-app
RUN npm install
RUN npm run build

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /harmonyserver/harmony-server .
COPY --from=builder /harmonyserver/.env .
COPY --from=harmonyclientbuilder /app/harmony-app/build ./static

EXPOSE 2288

CMD ["./harmony-server"]