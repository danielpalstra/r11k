SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=r11k

VERSION=$(shell git rev-parse HEAD)
BUILD_TIME=`date +%FT%T%z`

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	go build -o bin/${BINARY} main.go
