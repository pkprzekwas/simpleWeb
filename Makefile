SHELL := /bin/bash

TARGET := $(shell echo $${PWD\#\#*/})
.DEFAULT_GOAL: $(TARGET)

VERSION ?= 1.0.0
BUILD := $(shell git rev-parse HEAD)
APP := simple

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: all build clean fmt simplify

$(TARGET): $(SRC)
	@go build $(LDFLAGS) -o $(TARGET) -i cmd/$(APP)/main.go

build: $(TARGET)
	@true

clean:
	@rm -f $(TARGET)

fmt:
	@gofmt -l -w $(SRC)

simplify:
	@gofmt -s -l -w $(SRC)

db:
	docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=simple postgres:11
