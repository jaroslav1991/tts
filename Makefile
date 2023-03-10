PROJECT_NAME=cli

# go tool dist list
WINDOWS=windows/386 windows/amd64 windows/arm
DARWIN=darwin/amd64 darwin/arm64
LINUX=linux/386 linux/amd64 linux/arm linux/arm64
PLATFORMS=$(WINDOWS) $(LINUX) $(DARWIN)

.PHONY: build-all
build-all: $(PLATFORMS)

$(WINDOWS): export EXT=.exe

$(PLATFORMS): split=$(subst /, ,$@)
$(PLATFORMS): export OS=$(word 1,$(split))
$(PLATFORMS): export ARCH=$(word 2,$(split))
$(PLATFORMS):
	@$(MAKE) build

build:
	env GOOS=$(OS) GOARCH=$(ARCH) go build -o bin/$(PROJECT_NAME)-$(OS)-$(ARCH)$(EXT) cmd/cli/main.go

run: build-all

.PHONY: start-mock
start-mock:
	go run cmd/mock/main.go

.PHONY: send-test-event
send-test-event:
	go run ./cmd/cli/main.go -d '{"pluginType":"jetbrains","pluginVersion":"1.0.0","cliType":"macos","cliVersion":"2.1.0","deviceName":"vasyamac","events":[{"uid":"3607bbe0-2c9a-4c51-b636-5e6a7db8b574","createdAt":"2022-01-1114:23:01","type":"modifyfile","project":"someproject","language":"golang","target":"./"}]}'

.PHONY: help
help:
	go run ./cmd/cli/main.go -h
