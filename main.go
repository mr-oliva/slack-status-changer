package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/mitchellh/go-homedir"
)

type slack struct {
	WEB    string   `yaml:"internal_webapp_url"`
	Tokens []string `yaml:"tokens"`
}

func main() {
	slack := slack{}
	homedir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile(homedir + "/.slack.yml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &slack)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{}
	resp, err := client.Get(slack.WEB)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == 200 {
		slack.sendStatus("office")
	} else {
		slack.sendStatus("house")
	}
}

func (s *slack) sendStatus(status string) {
	endpoint, err := url.Parse("https://slack.com/api/users.profile.set")
	if err != nil {
		log.Fatal(err)
	}
	q := endpoint.Query()
	q.Add("profile", fmt.Sprintf(`{"status_emoji":":%s:"}`, status))
	for _, token := range s.Tokens {
		q.Add("token", token)
		tmpEndpoint := endpoint
		tmpEndpoint.RawQuery = q.Encode()
		_, err := http.Post(tmpEndpoint.String(), "application/json", nil)
		if err != nil {
			log.Fatal(err)
		}
		q.Del("token")
	}

}
