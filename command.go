package changer

type SlackClient interface {
	SendStatus(status string) error
}

type Command struct {
	InternalURL string
	SlackClient SlackClient
}

func (c *Command) Run() error {
	return nil
}
