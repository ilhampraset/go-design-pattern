package strategy

type Email struct {
}

func (e Email) Send() string {
	return "send mail"
}
