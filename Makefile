
.PHONY: all build test

BINARY_NAME=stroll

GO_SOURCES = $(shell find . -type f -name '*.go') go.mod

all: build

build: $(BINARY_NAME)

$(BINARY_NAME): $(GO_SOURCES)
	go build -o $(BINARY_NAME) *.go

test: $(GRAMMAR_FILE)
	go test -v -count=1 ./...