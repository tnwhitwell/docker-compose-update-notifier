package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/tnwhitwell/docker-compose-update-notifier/microbadger"

	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
)

type ImageString string

type ComposeFile struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
}

type Service struct {
	Image ImageString `yaml:"image"`
}

func (i *ImageString) string() string {
	return string(*i)
}

func (i *ImageString) parts() ([]string, error) {
	p := strings.Split(i.string(), ":")
	if len(p) != 2 {
		return []string{}, fmt.Errorf("Could not split image field")
	}
	return p, nil
}

func (i *ImageString) tag() (string, error) {
	parts, err := i.parts()
	if err != nil {
		return "", err
	}
	return parts[1], nil
}

func (i *ImageString) name() (string, error) {
	parts, err := i.parts()
	if err != nil {
		return "", err
	}
	return parts[0], nil
}

func (c *ComposeFile) parse(filePath *string) (*ComposeFile, error) {

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

func (s *Service) getLatest() (latestTag string, err error) {
	n, err := s.Image.name()
	if err != nil {
		return "", err
	}

	response, err := microbadger.GetImage(n)
	if err != nil {
		return "", err
	}

	return response.Versions[0].Tags[0].Tag, nil
}

func main() {
	var (
		// debug    = kingpin.Flag("debug", "Enable debug mode.").Bool()
		filePath = kingpin.Arg("filepath", "Number of packets to send").Required().ExistingFile()
	)
	kingpin.Version("0.0.1")
	kingpin.Parse()

	var composeFile ComposeFile

	composeFile.parse(filePath)

	for _, service := range composeFile.Services {
		latestTag, err := service.getLatest()
		if err != nil {
			log.Fatalf("error getting latest tag: %v", err)
		}
		currentTag, err := service.Image.tag()
		if err != nil {
			log.Fatalf("error getting current tag: %v", err)
		}
		if latestTag != currentTag {
			fmt.Println(latestTag)
		} else {
			fmt.Println(currentTag)
		}
	}

	// fmt.Printf("%+v", composeFile)
}
