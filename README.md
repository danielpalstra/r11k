# Description

Utility that servers as a wrapper around r10k and generates puppetfiles based
 YAML input.

## Build

This project uses `go-bindata` to embed go templates in the binary.
Install go-bindata by running

```shell
  go get -u github.com/jteeuwen/go-bindata/...
```

Run `make` and watch the magic. Binary will be put in bin/

## Usage

Run `bin/r11k --help` and a world will open.

## Roadmap

Perform git magic in r10k folder

*   Embed template in go binary
*   Add support for puppet module development
