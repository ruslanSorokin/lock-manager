.SILENT: tools.install
tools.install: tools.download
	@go mod download
	@cd docs && $(MAKE) --no-print-directory install-tools

_lint_vet:
	@(cd cmd && go vet ./...)
	@(cd pkg && go vet ./...)
	@(cd internal && go vet ./...)

_lint_imports:
	@goimports-reviser cmd pkg internal tools

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