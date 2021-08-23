package usecase

import (
	"context"
	"mygo/http/oapi"
)

// Signup ユーザ登録時に利用
func (u *UseCase) Signup(ctx context.Context, p *oapi.Signup) error {

	ac := u.db.Auth.Create().SetEmail(p.Email).SetPassword(p.Password)
	a, err := ac.Save(ctx)
	if err != nil {
		return err
	}

	uc := u.db.User.Create()
	uc.SetName(p.Name)
	uc.SetAuth(a)

	_, err = uc.Save(ctx)
	if err != nil {
		return err
	}

	return u.mailer.Send(p.Email, "Hello")
}
