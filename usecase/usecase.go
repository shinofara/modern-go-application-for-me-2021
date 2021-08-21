package usecase

import (
	"go.uber.org/dig"
	"mygo/ent"
	"mygo/infrastructure/mailer"
)

// UseCase usecaseに必要な設定や処理
type UseCase struct {
	DB     *ent.Client
	Mailer mailer.MailerInterface
}

// NewUseCase returns a new *UseCase
func NewUseCase(p struct{
	dig.In

	DB     *ent.Client
	Mailer mailer.MailerInterface
}) *UseCase {
	return &UseCase{
		DB:     p.DB,
		Mailer: p.Mailer,
	}
}