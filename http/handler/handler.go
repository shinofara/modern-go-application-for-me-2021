package handler

import (
	"github.com/shinofara/example-go-2021/ent"
	"github.com/shinofara/example-go-2021/infrastructure/mailer"
	"github.com/shinofara/example-go-2021/repository"
	"github.com/shinofara/example-go-2021/usecase"

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
