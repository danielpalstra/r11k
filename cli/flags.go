package cli

import "github.com/codegangsta/cli"

var (
	flEnvironment = cli.StringFlag{
		Name:  "environment, env",
		Usage: "Environment to generate the puppetfile for",
		Value: "all",
	}
	flInputYAML = cli.StringFlag{
		Name:  "input, in, i",
		Usage: "YAML file to use as input. ",
	}
	flOutputDir = cli.StringFlag{
		Name:  "output, out, o",
		Usage: "Directory to put output into",
		Value: "r10k",
	}
)
