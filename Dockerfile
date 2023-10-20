# syntax = docker/dockerfile:1-experimental

# Golang-Build container
FROM golang:1.21.0-alpine3.18 as builder

RUN apk add --no-cache \
  make=4.4.1-r1 \
  git=2.40.1-r0 \
  bash=5.2.15-r5 \
  build-base=0.5-r3 && \
  mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=${GOMODCACHE} \
  go mod download

COPY internal internal
COPY cmd cmd
COPY Makefile Makefile

RUN make app.build

# Distribution container
FROM alpine:3.18

LABEL maintainer="Ruslan Sorokin strawberryladder@gmail.com"

RUN apk add --no-cache \
  tzdata=2023c-r1 && \
  mkdir /app

WORKDIR /app

COPY config config

COPY --from=builder /app/main .

ENTRYPOINT [ "./main" ]

CMD [ "-config", "default.testing.yaml" ]
