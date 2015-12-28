package main

import (
	"github.com/danielpalstra/r11k/cli"
	_ "github.com/danielpalstra/r11k/puppetfile"
)

func main() {
	cli.Run()
}
