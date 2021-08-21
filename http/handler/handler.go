package handler

import (
	"mygo/ent"
	"mygo/interfaces"
)

type Handler struct {
	DB *ent.Client
	Mailer interfaces.MailerInterface
}

func NewHandler(db *ent.Client, mailer interfaces.MailerInterface) Handler {
	return Handler{
		DB: db,
		Mailer: mailer,
	}
}