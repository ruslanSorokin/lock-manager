install-gotask:
	go install -v github.com/go-task/task/v3/cmd/task@latest

.SILENT: app

app.unit-test:
	@go test ./... -count=1 -v -short

app.test:
	@go test ./... -count=1 -v

APP_ENTRYPOINT = ./cmd/lock-manager

app.build:
	@go build -v -o main $(APP_ENTRYPOINT)

app.run:
	@go run $(APP_ENTRYPOINT)
