package usecase

import (
	"context"
	"mygo/http/oapi"
)

// Signup ユーザ登録時に利用
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
