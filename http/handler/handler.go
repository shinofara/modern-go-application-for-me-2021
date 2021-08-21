package handler

import (
	"go.uber.org/dig"
	"mygo/ent"
	"mygo/infrastructure/mailer"
	"mygo/usecase"
)

type Handler struct {
	DB     *ent.Client
	Mailer mailer.MailerInterface
	UseCase *usecase.UseCase
}

func NewHandler(p struct{
	dig.In

	DB     *ent.Client
	Mailer mailer.MailerInterface
	UseCase *usecase.UseCase
}) Handler {
	return Handler{
		DB:     p.DB,
		Mailer: p.Mailer,
		UseCase: p.UseCase,
	}
}
