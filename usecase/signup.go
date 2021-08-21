package usecase

import (
	"context"
	"mygo/ent"
	oapi "mygo/http/openapi"
	"mygo/infrastructure/mailer"
)

type ServiceLocator struct {
	DB     *ent.Client
	Mailer mailer.MailerInterface
}

func Signup(ctx context.Context, sl *ServiceLocator, p *oapi.Signup) error {
	uc := sl.DB.User.Create()
	uc.SetName(p.Name)

	u, err := uc.Save(ctx)
	if err != nil {
		return err
	}

	ac := sl.DB.Auth.Create().SetEmail(p.Email).SetPassword(p.Password).SetUser(u)
	_, err = ac.Save(ctx)
	if err != nil {
		return err
	}

	return sl.Mailer.Send(p.Email, "Hello")
}
