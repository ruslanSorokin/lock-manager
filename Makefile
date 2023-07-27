.SILENT: docs.generate
docs.generate:
	@cd docs && $(MAKE) --no-print-directory generate-uml

.SILENT: tools.install
tools.install:
	@go mod download
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %
	@cd docs && $(MAKE) --no-print-directory install-tools

_lint_vet:
	@(cd cmd && go vet ./...)
	@(cd internal && go vet ./...)

_lint_imports:
	@goimports-reviser cmd internal

_lint_golangci:
	@golangci-lint run

.SILENT: lint
lint: _lint_vet _lint_imports _lint_golangci

.SILENT: short-test.run
short-test.run:
	@go test ./... -count=1 -v -short

.SILENT: test.run
test.run:
	@go test ./... -count=1 -v

.SILENT: build
build:
	@go build -v -o main ./cmd/lock-manager

.SILENT: run
run:
	@go run ./cmd/lock-manager

.SILENT: docker.build
docker.build:
	@docker build --file Dockerfile --tag lock-manager .

.SILENT: docker.up
docker.up:
	@cd deploy && docker-compose -f docker-compose.yaml -f infra/redis/docker-compose.override.yaml up --build lock-manager -d

.SILENT: docker.down
docker.down:
	@cd deploy && docker-compose -f docker-compose.yaml -f infra/redis/docker-compose.override.yaml down
