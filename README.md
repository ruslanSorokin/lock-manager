<h1><a id="summary" class="anchor" aria-hidden="true"></a>Interprocess Communication Lock Manager</h1>

<h2><a id="summary" class="anchor" aria-hidden="true"></a>TLDR</h2>

Provides functionality to acquire & release resources over gRPC & HTTP.

[`API`](https://github.com/ruslanSorokin/lock-manager-api)

<h2><a id="summary" class="anchor" aria-hidden="true"></a>Tech Stack</h2>

<h3><a id="summary" class="anchor" aria-hidden="true"></a>Golang</h3>

- [`google/wire`](https://github.com/google/wire) for compile-time DI
- [`oklog/run`](https://github.com/oklog/run) for graceful structured
  concurrency
- [`ilyakaznacheev/cleanenv`](https://github.com/ilyakaznacheev/cleanenv) for YAML/TOML/JSON/ENV configuration
- [`gofiber/fiber`](https://github.com/gofiber/fiber): HTTP transport
- [`grpc/grpc-go`](https://github.com/grpc/grpc-go): GRPC transport
- [`go-playground/validator`](https://github.com/go-playground/validator): validation
- [`redis/redis`](https://github.com/redis/redis) as in-memory key-value database
- [`redis/go-redis`](https://github.com/redis/go-redis) as redis driver
- [`ory/dockertest`](https://github.com/ory/dockertest): container
  orchestration in integration tests
- [`vektra/mockery`](https://github.com/vektra/mockery): mocks for table-driven unit tests
- [`stretchr/testify`](https://github.com/stretchr/testify): suites for interface-driven integration tests
- [`prometheus/client_golang`](https://github.com/prometheus/client_golang): metrics
- [`go-logr/logr`](https://github.com/go-logr/logr) with [`rs/zerolog`](https://github.com/rs/zerolog): logging

---

<h3><a id="summary" class="anchor" aria-hidden="true"></a>General</h3>

- [`pre-commit/pre-commit`](https://github.com/pre-commit/pre-commit): git hooks management
- [`hadolint/hadolint`](https://github.com/hadolint/hadolint): Dockerfile linting
- [`golangci/golangci-lint`](https://github.com/golangci/golangci-lint) as Golang linter-runner
- [`go-task/task`](https://github.com/go-task/task) as task runner(Makefile alternative)

---

<h2><a id="summary" class="anchor" aria-hidden="true"></a>How to run</h2>

- `make install-gotask` to install go-task
- `task deploy:up` to deploy application & all needed infrastructure locally
- `task --list` to see all targets with description
