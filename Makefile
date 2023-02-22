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
	go build -o nil/$(PROJECT_NAME).linux-arm64 cmd/cli/main.go

	set GOOS=windows
	set GOARCH=amd64
	go build -o bin/$(PROJECT_NAME).win-amd64.exe cmd/cli/main.go

	set GOOS=windows
	set GOARCH=arm64
	go build -o bin/$(PROJECT_NAME).win-arm64.exe cmd/cli/main.go


run: build_all


