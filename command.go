package changer

import "net/http"

//go:generate counterfeiter -o changerFakes/slack.go . SlackClient
type SlackClient interface {
	SendStatus(status string) error
}

type Command struct {
	InternalURL string
	SlackClient SlackClient
	HttpClient  *http.Client
}

func (c *Command) Run() error {
	res, err := c.HttpClient.Get(c.InternalURL)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		if err := c.SlackClient.SendStatus("house"); err != nil {
			return err
		}
		return nil
	}
	if err := c.SlackClient.SendStatus("office"); err != nil {
		return err
	}
	return nil
}
