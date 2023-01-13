package strategy

type MessageStrategyMap map[string]NotificationSender

func (m MessageStrategyMap) Register(key string, value NotificationSender) {
	m[key] = value
}

func (m MessageStrategyMap) Send() string {
	return m.Send()
}
