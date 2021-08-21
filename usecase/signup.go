package usecase

import (
	"context"
	"go.uber.org/dig"
	"mygo/ent"
	oapi "mygo/http/openapi"
	"mygo/infrastructure/mailer"
)

type UseCase struct {
	DB     *ent.Client
	Mailer mailer.MailerInterface
}

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

func (u *UseCase) Signup(ctx context.Context, p *oapi.Signup) error {
	uc := u.DB.User.Create()
	uc.SetName(p.Name)

	user, err := uc.Save(ctx)
	if err != nil {
		return err
	}

	ac := u.DB.Auth.Create().SetEmail(p.Email).SetPassword(p.Password).SetUser(user)
	_, err = ac.Save(ctx)
	if err != nil {
		return err
	}

	return u.Mailer.Send(p.Email, "Hello")
}
