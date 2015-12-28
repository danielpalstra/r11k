package puppetfile

import (
	"bytes"
	"io/ioutil"
	"os"

	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v2"
)

// Puppetfile represents the yaml presentation from a Puppetfile
type Puppetfile struct {
	Modules map[string]Mod `yaml:"puppetfile"`
}

// Mod represents a module definition
type Mod struct {
	GitRemote string `yaml:"git"`
	Ref       string `yaml:"ref,omitempty"`
	Tag       string `yaml:"tag,omitempty"`
}

// LoadYaml  opens the provided YAML file and returns a struct
func LoadYaml(fileName string) (p Puppetfile, err error) {
	filename, _ := filepath.Abs(fileName)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(yamlFile, &p)
	if err != nil {
		return
	}
	return
}

// WriteToFile writes the buffered stream to file based on output var
func WriteToFile(output string, b bytes.Buffer) error {
	f, err := os.Create(output)
	b.WriteTo(f)
	return err
}

// BuildPuppetFile parses a template that creates the puppet file
func BuildPuppetFile(p Puppetfile) (b bytes.Buffer, err error) {
	// TODO package in binary or ma
	// https://github.com/jteeuwen/go-bindata
	t, err := template.ParseFiles("puppetfile/Puppetfile.tpl")
	t.Execute(&b, p)

	return
}
