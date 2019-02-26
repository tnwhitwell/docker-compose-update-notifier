package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
)

type ComposeFile struct {
	Version  string `yaml:"version"`
	Services map[string]Service
}

type Service struct {
	Image string
}

func (c *ComposeFile) getConfig(filePath *string) (*ComposeFile, error) {

	yamlFile, err := ioutil.ReadFile(*filePath)
	if err != nil {
		return c, fmt.Errorf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return c, fmt.Errorf("Unmarshal: %v", err)
	}

	return c, nil
}

func main() {
	var (
		// debug    = kingpin.Flag("debug", "Enable debug mode.").Bool()
		filePath = kingpin.Arg("filepath", "Number of packets to send").Required().ExistingFile()
	)
	kingpin.Version("0.0.1")
	kingpin.Parse()

	var composeFile ComposeFile

	composeFile.getConfig(filePath)

	fmt.Printf("%+v", composeFile)
}
