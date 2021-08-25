package usecase

import (
	"github.com/shinofara/example-go-2021/ent"
	"github.com/shinofara/example-go-2021/infrastructure/mailer"
	"github.com/shinofara/example-go-2021/repository"

	"go.uber.org/dig"
)

// UseCase usecaseに必要な設定や処理
type UseCase struct {
	db         *ent.Client
	mailer     mailer.MailerInterface
	repository *repository.Repository
}

// NewUseCase returns a new *UseCase
func NewUseCase(p struct {
	dig.In

	DB         *ent.Client
	Mailer     mailer.MailerInterface
	Repository *repository.Repository
}) *UseCase {
	return &UseCase{
		db:         p.DB,
		mailer:     p.Mailer,
		repository: p.Repository,
	}
}
