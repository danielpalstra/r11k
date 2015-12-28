package cli

import "github.com/codegangsta/cli"

var (
	commands = []cli.Command{
		{
			Name:      "generate",
			ShortName: "g",
			Usage:     "Generate puppetfile for the environments",
			Action:    generate,
			Flags:     []cli.Flag{flEnvironment, flInputYAML, flOutputDir},
		},
	}
)
