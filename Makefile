# Buildfile for r11k Go app

.PHONY: build

# VERSION should be updated by hand at each release
VERSION			:= 0.9.0
BUILD_TIME	:= `date +%FT%T%z`

# Get architecture from shell
GOOS				:= $(shell go env GOOS)
GOARCH			:= $(shell go env GOARCH)

# Set build props
BINARY			:= r11k
LDFLAGS			:= -ldflags "-X github.com/danielpalstra/r11k/version.VERSION=${VERSION} -X github.com/danielpalstra/r11k/version.GITCOMMIT=`git rev-parse --short HEAD`"
GOBUILD			:= GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS)
BIN_DIR			:= bin

default: build

build:
	go generate
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY)-$(GOOS)-$(GOARCH)
	# env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY}-${GOOS}-${GOARCH}
