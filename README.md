# Interprocess Communication Lock Manager

## TLDR

Provides functionality to acquire & release resources over gRPC & HTTP.

[`API`](https://github.com/ruslanSorokin/lock-manager-api)

## Tech Stack

### Golang

- [`google/wire`](https://github.com/google/wire) for compile-time DI
- [`oklog/run`](https://github.com/oklog/run) for graceful structured
  concurrency
- [`ilyakaznacheev/cleanenv`](https://github.com/ilyakaznacheev/cleanenv) for YAML/TOML/JSON/ENV configuration
- [`gofiber/fiber`](https://github.com/gofiber): HTTP transport
- [`grpc/grpc-go`](https://github.com/grpc/grpc-go): GRPC transport
- [`go-playground/validator`](https://github.com/go-playground/validator): validation
- [`redis/redis`](https://github.com/redis/redis) as in-memory key-value database
- [`redis/go-redis`](https://github.com/redis/go-redis) as redis driver
- [`ory/dockertest`](https://github.com/ory/dockertest) for container
  orchestration in integration tests
- [`vektra/mockery`](https://github.com/vektra/mockery): mocks for table-driven unit tests
- [`stretchr/testify`](https://github.com/stretchr/testify): suites for interface-driven integration tests
- [`prometheus/client_golang`](https://github.com/prometheus/client_golang): metrics
- [`go-logr/logr`](https://github.com/go-logr/logr) with [`rs/zerolog`](https://github.com/rs/zerolog) for logging

---

### General

- [`pre-commit/pre-commit`](https://github.com/pre-commit/pre-commit): git hooks management
- [`hadolint/hadolint`](https://github.com/hadolint/hadolint): Dockerfile linter
- [`golangci/golangci-lint`](https://github.com/golangci/golangci-lint): Golang linter-runner
- [`go-task/task`](https://github.com/go-task/task): task runner

---

## How to run

- Install go-task with `make install-gotask`
- `task deploy:up` to deploy application & all needed infrastructure locally
- `task --list` to see all targets with description
