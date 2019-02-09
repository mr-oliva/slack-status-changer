package changer

import (
	"fmt"
	"net/http"
	"net/url"
)

type Slack struct {
	Tokens   []string `yaml:"tokens"`
}

func (s *Slack) SendStatus(status string) error {
	endpoint, err := url.Parse("https://slack.com/api/users.profile.set")
	if err != nil {
		return err
	}
	q := endpoint.Query()
	q.Add("profile", fmt.Sprintf(`{"status_emoji":":%s:"}`, status))
	for _, token := range s.Tokens {
		q.Add("token", token)
		tmpEndpoint := endpoint
		tmpEndpoint.RawQuery = q.Encode()
		_, err := http.Post(tmpEndpoint.String(), "application/json", nil)
		if err != nil {
			return err
		}
		q.Del("token")
	}
	return nil
}
