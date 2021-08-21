package interfaces

type MailerInterface interface {
	Send(email, title string) error
}
