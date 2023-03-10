PROJECT_NAME=cli
BUILD_DIR=./bin

# go tool dist list
WINDOWS=windows/386 windows/amd64 windows/arm
DARWIN=darwin/amd64 darwin/arm64
LINUX=linux/386 linux/amd64 linux/arm linux/arm64
#PLATFORMS=$(WINDOWS) $(LINUX) $(DARWIN)
PLATFORMS=darwin/arm64

run: build-all

.PHONY: build-all
build-all: $(PLATFORMS)

$(WINDOWS): EXT=.exe
$(PLATFORMS): split=$(subst /, ,$@)
$(PLATFORMS): OS=$(word 1,$(split))
$(PLATFORMS): ARCH=$(word 2,$(split))
$(PLATFORMS): ARTIFACT_NAME=$(PROJECT_NAME)-$(OS)-$(ARCH)$(EXT)
$(PLATFORMS):
	env GOOS=$(OS) GOARCH=$(ARCH) go build -o $(BUILD_DIR)/$(ARTIFACT_NAME) cmd/cli/main.go

.PHONY: zip-artifacts
zip-artifacts: $(foreach f,$(wildcard $(BUILD_DIR)/*),$(f).zip)

$(BUILD_DIR)/%.zip:
	@cd $(BUILD_DIR) && zip $*.zip $*

.PHONY: start-mock
start-mock:
	go run cmd/mock/main.go

.PHONY: send-test-event
send-test-event:
	go run ./cmd/cli/main.go -d '{"pluginType":"jetbrains","pluginVersion":"1.0.0","cliType":"macos","cliVersion":"2.1.0","deviceName":"vasyamac","events":[{"uid":"3607bbe0-2c9a-4c51-b636-5e6a7db8b574","createdAt":"2022-01-1114:23:01","type":"modifyfile","project":"someproject","language":"golang","target":"./"}]}'

.PHONY: help
help:
	go run ./cmd/cli/main.go -h
