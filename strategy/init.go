package strategy

var Message = make(MessageStrategyMap)

func InitMessageStrategy() {
	MessageRegister()

}

func MessageRegister() {
	Message.Register("1", &Whatsapp{})
	Message.Register("2", &Email{})
	Message.Register("3", &Slack{})
	Message.Register("4", &Telegram{})
}
