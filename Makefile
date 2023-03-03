PROJECT_NAME=cli

.PHONY: build_all
build_all:
	set GOOS=darwin
	set GOARCH=amd64
	go build -o bin/$(PROJECT_NAME).mac-amd64 cmd/cli/main.go

	set GOOS=darwin
	set GOARCH=arm64
	go build -o bin/$(PROJECT_NAME).mac-arm64 cmd/cli/main.go

	set GOOS=linux
	set GOARCH=amd64
	go build -o bin/$(PROJECT_NAME).linux-amd64 cmd/cli/main.go

	set GOOS=linux
	set GOARCH=arm64
	go build -o bin/$(PROJECT_NAME).linux-arm64 cmd/cli/main.go

	set GOOS=windows
	set GOARCH=amd64
	go build -o bin/$(PROJECT_NAME).win-amd64.exe cmd/cli/main.go

	set GOOS=windows
	set GOARCH=arm64
	go build -o bin/$(PROJECT_NAME).win-arm64.exe cmd/cli/main.go


run: build_all


.PHONY: start-mock
start-mock:
	go run cmd/mock/main.go

.PHONY: send-test-event
send-test-event:
	go run ./cmd/cli/main.go -d '{"pluginType":"jetbrains","pluginVersion":"1.0.0","cliType":"macos","cliVersion":"2.1.0","deviceName":"vasyamac","events":[{"uid":"3607bbe0-2c9a-4c51-b636-5e6a7db8b574","createdAt":"2022-01-1114:23:01","type":"modifyfile","project":"someproject","language":"golang","target":"./"}]}'

.PHONY: help
help:
	go run ./cmd/cli/main.go -h

