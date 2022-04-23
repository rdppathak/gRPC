#
# Author: rdppathak@gmail.com
#
# Build & tools related target for project

GO_CMD := go
GO_BUILD := $(GO_CMD) build
GO_CLEAN := $(GO_CMD) clean
GO_TEST := $(GO_CMD) test
GO_FMT := $(GO_CMD) fmt

BUILD_DIR := build
build-dir:
	mkdir -p $(BUILD_DIR)

build-server: build-dir
	@echo "Building server binary..."
	$(GO_BUILD) -o $(BUILD_DIR)/server cmd/server/main.go
.PHONY: build-server

clean:
	rm -rf $(BUILD_DIR)
.PHONY: clean

