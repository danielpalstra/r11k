package cli

import (
	"log"

	pf "github.com/danielpalstra/r11k/puppetfile"

	"github.com/codegangsta/cli"
)

func generate(c *cli.Context) {
	// Logic for handling creation
	f := c.String("input")
	if f == "" {
		log.Fatal("--input flag is mandatory")
	}

	// Define the name for the puppetfile
	o := c.String("output") + "/Puppetfile"
	// if o == "" {
	// 	log.Fatal("--output flag is mandatory")
	// }

	p, err := pf.LoadYaml(f)
	if err != nil {
		log.Fatalf("Error opening YAML file %v, error:  %v\n", f, err)
	}

	// Build the puppet file from YAML
	b, err := pf.BuildPuppetFile(p)
	if err != nil {
		log.Fatalf("Error generating puppet file from struct: %v \n", err)
	}

	err = pf.WriteToFile(o, b)
	if err != nil {
		log.Fatalf("Error creating puppetfile: %v \n", err)
	}

	log.Printf("Puppet file written to: %v \n", o)

	// Write output
	// TODO write to file instead

	// b.WriteTo(os.Stdout)

}
