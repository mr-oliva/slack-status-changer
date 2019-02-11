package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bookun/slack-status-changer"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

type Manifest struct {
	WEB    string   `yaml:"internal_webapp_url"`
	Tokens []string `yaml:"tokens"`
}

func main() {
	homedir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile(homedir + "/.slack.yml")
	if err != nil {
		log.Fatal(err)
	}
	var manifest Manifest
	err = yaml.Unmarshal(data, &manifest)
	if err != nil {
		log.Fatal(err)
	}
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	config := changer.Config{
		HTTPClient: httpClient,
	}
	slackClient := changer.NewSlackClient(config, manifest.Tokens)
	command := changer.Command{
		InternalURL: manifest.WEB,
		SlackClient: slackClient,
		HttpClient:  httpClient,
	}
	if err := command.Run(); err != nil {
		log.Fatal(err)
	}
	return
}
