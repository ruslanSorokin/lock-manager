###############################################################################

.SILENT: docs

docs.generate:
	@cd docs && $(MAKE) --no-print-directory generate-uml

###############################################################################

.SILENT: tools

GO_TOOLS_FILE = tools/tools.go

tools.install:
	@go mod download
	@cat  $(GO_TOOLS_FILE) | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %
	@cd docs && $(MAKE) --no-print-directory install-tools

_lint_vet:
	@(cd cmd && go vet ./...)
	@(cd internal && go vet ./...)

_lint_imports:
	@goimports-reviser cmd internal

_lint_golangci:
	@golangci-lint run

###############################################################################

.SILENT: lint

lint: _lint_vet _lint_imports _lint_golangci

###############################################################################

.SILENT: app

app.unit-test:
	@go test ./... -count=1 -v -short

app.test:
	@go test ./... -count=1 -v

APP_ENTRYPOINT = ./cmd/lock-manager

app.build:
	@go build -v -o main $(APP_ENTRYPOINT)

app.run:
	@go run $(APP_ENTRYPOINT) -config local


###############################################################################

.SILENT: docker

DOCKER_TAG = lock-manager

docker.build:
	@docker build --file Dockerfile --tag $(DOCKER_TAG) .

DOCKER_COMPOSE_ROOT = deploy/docker

docker.up:
	@cd $(DOCKER_COMPOSE_ROOT) && docker-compose -f docker-compose.yaml -f infra/storage/redis/docker-compose.override.yaml up --build $(DOCKER_TAG) -d

docker.down:
	@cd $(DOCKER_COMPOSE_ROOT) && docker-compose -f docker-compose.yaml -f infra/storage/redis/docker-compose.override.yaml down

docker.restart: docker.down docker.up
