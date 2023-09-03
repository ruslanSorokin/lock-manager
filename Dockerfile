# syntax = docker/dockerfile:1-experimental

# Golang-Build container
FROM golang:1.21.0-alpine3.18 as builder

RUN apk add git make bash build-base && \
  mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=${GOMODCACHE} \
  go mod download

COPY . .

RUN make app.build

# Distribution container
FROM alpine:3.18

LABEL maintainer="Ruslan Sorokin strawberryladder@gmail.com"

RUN apk add --no-cache tzdata && \
  mkdir /app

WORKDIR /app

COPY config/ config/

COPY --from=builder /app/main .

ENTRYPOINT [ "./main" ]
