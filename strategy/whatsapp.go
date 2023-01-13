package strategy

type Whatsapp struct {
}

func (w Whatsapp) Send() string {
	return "send whatsapp"
}
