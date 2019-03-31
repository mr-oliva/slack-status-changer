package changer

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Config struct {
	Endpoint   string
	HTTPClient *http.Client
}

func DefaultConfig() Config {
	return Config{
		Endpoint:   "https://slack.com/api",
		HTTPClient: http.DefaultClient,
	}
}

type Slack struct {
	Config Config
	Tokens []string
}

type Response struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

func NewSlackClient(config Config, tokens []string) *Slack {
	defConfig := DefaultConfig()
	if config.Endpoint == "" {
		config.Endpoint = defConfig.Endpoint
	}
	if config.HTTPClient == nil {
		config.HTTPClient = defConfig.HTTPClient
	}
	return &Slack{
		Config: config,
		Tokens: tokens,
	}
}

func (s *Slack) SendStatus(status string) error {
	u := fmt.Sprintf("%s/users.profile.set", s.Config.Endpoint)
	values := url.Values{}
	values.Set("profile", fmt.Sprintf(`{"status_emoji":":%s:"}`, status))
	for _, token := range s.Tokens {
		req, err := http.NewRequest("POST", u, strings.NewReader(values.Encode()))
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		slackResponse, err := s.doPost(req)
		if err != nil {
			return err
		}
		if !slackResponse.Ok {
			return errors.New(slackResponse.Error)
		}
	}
	return nil
}

func (s *Slack) doPost(req *http.Request) (*Response, error) {
	resp, err := s.Config.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var response Response
	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}
	return &response, err
}
