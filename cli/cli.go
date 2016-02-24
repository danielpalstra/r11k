package cli

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/danielpalstra/r11k/version"
)

// Run the Swarm CLI.
func Run() {
	app := cli.NewApp()
	app.Name = "r11k"
	app.Usage = "r10k Puppetfile on steroids"
	app.Version = version.VERSION + " (" + version.GITCOMMIT + ")"

	app.Author = ""
	app.Email = ""

	app.Commands = commands

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
