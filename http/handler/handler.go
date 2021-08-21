package handler

import (
	"mygo/ent"
	"mygo/infrastructure/mailer"
)

type Handler struct {
	DB     *ent.Client
	Mailer mailer.MailerInterface
}

func NewHandler(db *ent.Client, mailer mailer.MailerInterface) Handler {
	return Handler{
		DB:     db,
		Mailer: mailer,
	}
}
