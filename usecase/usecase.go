package usecase

import (
	"go.uber.org/dig"
	"mygo/ent"
	"mygo/infrastructure/mailer"
	"mygo/repository"
)

// UseCase usecaseに必要な設定や処理
type UseCase struct {
	DB     *ent.Client
	Mailer mailer.MailerInterface
	Repository *repository.Repository
}

// NewUseCase returns a new *UseCase
func NewUseCase(p struct{
	dig.In

	DB     *ent.Client
	Mailer mailer.MailerInterface
	Repository *repository.Repository
}) *UseCase {
	return &UseCase{
		DB:     p.DB,
		Mailer: p.Mailer,
		Repository: p.Repository,
	}
}