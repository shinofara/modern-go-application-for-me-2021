package usecase

import (
	"context"
	"mygo/ent"
	oapi "mygo/http/openapi"
)

func Signup(ctx context.Context, db *ent.Client, p *oapi.Signup) error {
	u := db.User.Create()
	u.SetEmail(p.Email).SetName(p.Name).SetPassword(p.Password)
	if _, err := u.Save(ctx); err != nil {
		return err
	}
	return nil
}
