CURRENT_REVISION = $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS = "-s -w -X main.revision=$(CURRENT_REVISION)"

.PHONY: build
build:
	go build -o gtt -ldflags=$(BUILD_LDFLAGS)
