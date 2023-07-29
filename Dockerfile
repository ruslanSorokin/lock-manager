# syntax = docker/dockerfile:1-experimental

# Builder
FROM golang:1.20.6-alpine3.18 as builder

ARG GOMODCACHE
ARG GOCACHE

RUN apk add git make bash build-base && \
  mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=${GOMODCACHE} \
  go mod download

COPY . .

RUN --mount=type=cache,target=${GOCACHE} \
  make build

# Distribution
FROM alpine:3.18

LABEL maintainer="Ruslan Sorokin strawberryladder@gmail.com"

RUN apk add --no-cache tzdata && \
  mkdir /app

WORKDIR /app

EXPOSE 8082

COPY --from=builder /app/main .
COPY configs/ configs/

CMD ["./main"]
