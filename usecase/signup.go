package usecase

import (
	"context"
	"mygo/ent"
	oapi "mygo/http/openapi"
)

func Signup(ctx context.Context, db *ent.Client, p *oapi.Signup) error {
	uc := db.User.Create()
	uc.SetName(p.Name)

	u, err := uc.Save(ctx)
	if err != nil {
		return err
	}

	ac := db.Auth.Create().SetEmail(p.Email).SetPassword(p.Password).SetUser(u)
	_, err = ac.Save(ctx)

	return err
}
