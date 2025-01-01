
BINARY_NAME=myapp

SRC_DIR=./cmd
MIGRATE_DIR=./migrations

GO=go
GOFMT=gofmt
GOFLAGS=-v
LDFLAGS=-s -w

migrate:
	$(GO) run $(MIGRATE_DIR)
build:
	$(GO) build $(GOFLAGS) -o $(BINARY_NAME) $(SRC_DIR)
run:
	$(GO) run $(SRC_DIR)
test:
	$(GO) test $(GOFLAGS) ./...
deps:
	$(GO) mod tidy

all: test build
