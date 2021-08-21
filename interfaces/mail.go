package interfaces

type MailerInterface interface {
	Send(email, title string) error
}

func NewDummyMailer() MailerInterface{
	return &DummyMailer{}
}

type DummyMailer struct {}
func (DummyMailer) Send(email, title string) error {
	return nil
}