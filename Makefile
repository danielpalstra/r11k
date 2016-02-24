.PHONY: build

BINARY=r11k

# VERSION should be updated by hand at each release
VERSION=0.9.0
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-X github.com/danielpalstra/r11k/version.VERSION=${VERSION} -X github.com/danielpalstra/r11k/version.GITCOMMIT=`git rev-parse --short HEAD`"

default: build

build:
	go generate
	go build ${LDFLAGS} -o bin/${BINARY}
