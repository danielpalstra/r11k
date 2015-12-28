package cli

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

// Run the Swarm CLI.
func Run() {
	app := cli.NewApp()
	app.Name = "r10k-template"
	app.Usage = "A r10k template util"

	app.Author = ""
	app.Email = ""

	app.Commands = commands

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
