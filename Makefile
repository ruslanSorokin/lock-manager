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
	@goimports-reviser cmd internal tools

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
	@go build -v cmd/afterwork-backend/main.go
