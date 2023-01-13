package strategy

type Telegram struct {
}

func (e Telegram) Send() string {
	return "send telegram"
}
