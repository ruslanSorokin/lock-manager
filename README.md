# Interprocess Communication Lock Manager

### TLDR

Provides functionality to acquire & release resources over gRPC & HTTP.

[API repository](https://github.com/ruslanSorokin/lock-manager-api)

### Tech Stack

- [`google/wire`](https://github.com/google/wire) for compile-time wiring
- [`oklog/run`](https://github.com/oklog/run) for graceful structured
  concurrency
- [`ilyakaznacheev/cleanenv`](https://github.com/ilyakaznacheev/cleanenv) for `YAML/TOML/JSON/ENV` configuration
- [`gofiber/fiber`](https://github.com/gofiber) for `HTTP` transport
- [`grpc/grpc-go`](https://github.com/grpc/grpc-go) for `GRPC` transport
- [`go-playground/validator`](https://github.com/go-playground/validator) for validation
- [`redis/redis`](https://github.com/redis/redis) as in-memory key-value Database
- [`redis/go-redis`](https://github.com/redis/go-redis) for redis interaction
- [`ory/dockertest`](https://github.com/ory/dockertest) for container
  orchestration in Integration tests
- [`stretchr/testify`](https://github.com/stretchr/testify) for `table-driven` unit
  tests with mocks & for `interface-driven` integration tests with suites
- [`prometheus/client_golang`](https://github.com/prometheus/client_golang) for metrics
- [`go-logr/logr`](https://github.com/go-logr/logr) with [`rs/zerolog`](https://github.com/rs/zerolog) for logging

### Overall Project structure

```bash
├── .config # configs for linters, etc.
├── .github # github actions CI files
│  └── workflows
├── ci # helpers for CI, e.g. Taskfiles, Makefiles, etc.
├── cmd # entry points
│  └── lock-manager
├── config # yaml/toml/json/env configs
│  └── lock-manager
├── deploy # local deploy stuff
│  └── docker # docker-compose local deploy stuff
│     ├── grafana
│     └── prometheus
├── docs
│  ├── gen # documentation
│  └── src
├── internal
│  ├── lock-manager
│  │  ├── app # wiring
│  │  ├── handler # inbound adapters
│  │  │  ├── ifiber # http
│  │  │  │  ├── lock
│  │  │  │  ├── shared
│  │  │  │  └── unlock
│  │  │  └── igrpc # grpc
│  │  │     ├── lock
│  │  │     ├── shared
│  │  │     └── unlock
│  │  ├── ierror # error type
│  │  ├── ilog # some shared log tags
│  │  ├── metric # application
│  │  │  ├── iprom
│  │  │  └── mock
│  │  ├── model # domain models
│  │  ├── provider # outbound adapters
│  │  │  ├── mock
│  │  │  ├── storage
│  │  │  │  └── iredis
│  │  │  └── test
│  │  └── service # service
│  │     └── mock
│  └── pkg # boilerplate
│     ├── conn
│     │  └── redis
│     └── util
│        ├── app
│        │  ├── iprom
│        │  └── mock
│        ├── config
│        ├── fiber
│        ├── grpc
│        │  ├── iprom
│        │  └── mock
│        ├── http
│        └── prom
├── script # scripts for CI
└── tools # tools checked in with go.mod
```
