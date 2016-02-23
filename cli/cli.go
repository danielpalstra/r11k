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
	app.Name = "r10k-template"
	app.Usage = "A r10k template util"
	app.Version = version.VERSION + " (" + version.GITCOMMIT + ")"

	app.Author = ""
	app.Email = ""

	app.Commands = commands

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
