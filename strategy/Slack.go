package strategy

type Slack struct {
}

func (s Slack) Send() string {
	return "send slack"
}
