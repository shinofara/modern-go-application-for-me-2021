package handler

import (
	"mygo/ent"
	"mygo/infrastructure/mailer"
	"mygo/repository"
	"mygo/usecase"

	"go.uber.org/dig"
)

type Handler struct {
	db         *ent.Client
	mailer     mailer.MailerInterface
	useCase    *usecase.UseCase
	repository *repository.Repository
}

func NewHandler(p struct {
	dig.In

	DB         *ent.Client
	Mailer     mailer.MailerInterface
	UseCase    *usecase.UseCase
	Repository *repository.Repository
}) Handler {
	return Handler{
		db:         p.DB,
		mailer:     p.Mailer,
		useCase:    p.UseCase,
		repository: p.Repository,
	}
}
