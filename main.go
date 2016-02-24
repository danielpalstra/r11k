package main

import (
	"github.com/danielpalstra/r11k/cli"
	_ "github.com/danielpalstra/r11k/puppetfile"
)

//go:generate go-bindata -pkg puppetfile -o puppetfile/assets.go puppetfile/Puppetfile.tpl

func main() {
	cli.Run()
}
