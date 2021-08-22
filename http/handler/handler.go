package handler

import (
	"go.uber.org/dig"
	"mygo/ent"
	"mygo/infrastructure/mailer"
	"mygo/repository"
	"mygo/usecase"
)

type Handler struct {
	DB     *ent.Client
	Mailer mailer.MailerInterface
	UseCase *usecase.UseCase
	Repository *repository.Repository
}

func NewHandler(p struct{
	dig.In

	DB     *ent.Client
	Mailer mailer.MailerInterface
	UseCase *usecase.UseCase
	Repository *repository.Repository
}) Handler {
	return Handler{
		DB:     p.DB,
		Mailer: p.Mailer,
		UseCase: p.UseCase,
		Repository: p.Repository,
	}
}
