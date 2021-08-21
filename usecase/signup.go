package usecase

import (
	"context"
	"mygo/ent"
	oapi "mygo/http/openapi"
	"mygo/interfaces"
)

func Signup(ctx context.Context, db *ent.Client, mailer interfaces.MailerInterface,p *oapi.Signup) error {
	uc := db.User.Create()
	uc.SetName(p.Name)

	u, err := uc.Save(ctx)
	if err != nil {
		return err
	}

	ac := db.Auth.Create().SetEmail(p.Email).SetPassword(p.Password).SetUser(u)
	_, err = ac.Save(ctx)
	if err != nil {
		return err
	}

	return mailer.Send(p.Email, "Hello")
}
