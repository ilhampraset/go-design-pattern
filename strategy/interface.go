package strategy

type NotificationSender interface {
	Send() (string, err)
}
