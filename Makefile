# Makefile
export GO111MODULE=on

GO_CMD=go
GO_MOD_TIDY=$(GO_CMD) mod tidy
GO_GET=$(GO_CMD) get -v
GO_BUILD=$(GO_CMD) build
GO_BUILD_RACE=$(GO_CMD) build -race
GO_TEST=$(GO_CMD) test
GO_TEST_VERBOSE=$(GO_CMD) test -v
GO_TEST_COVER=$(GO_CMD) test -cover -count=1 -p=1
GO_INSTALL=$(GO_CMD) install -v

SERVER_BIN=last9
SERVER_BIN_DIR=.
SERVER_DIR=.
SERVER_MAIN=main.go

SOURCE_PKG_DIR= .
SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

all: dependencies build-server

dependencies:
	@echo "==> Installing dependencies ...";
	@$(GO_MOD_TIDY)

build-server:
	@echo "==> Building server ...";
	@GOOS=darwin $(GO_BUILD) -o $(SERVER_BIN) -ldflags "-w -s" $(SERVER_DIR)/$(SERVER_MAIN) || exit 1;
	@chmod 755 $(SERVER_BIN)

run: dependencies build-server
	./$(SERVER_BIN) config.json