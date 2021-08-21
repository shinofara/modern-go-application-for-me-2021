package interfaces

import "log"

type MailerInterface interface {
	Send(email, title string) error
}

func NewDummyMailer() MailerInterface{
	return &DummyMailer{}
}

type DummyMailer struct {}
func (DummyMailer) Send(email, title string) error {
	log.Println(email, title)
	return nil
}