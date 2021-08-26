package usecase

import (
	"context"
	"database/sql"

	"github.com/shinofara/example-go-2021/openapi"
)

// Signup ユーザ登録時に利用
func (u *UseCase) Signup(ctx context.Context, p *openapi.SignupRequest) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			if err != sql.ErrTxDone {
				panic(err)
			}

		}
	}()

	ac := tx.Auth.Create().SetEmail(p.Email).SetPassword(p.Password)
	a, err := ac.Save(ctx)
	if err != nil {
		return err
	}

	uc := tx.User.Create()
	uc.SetName(p.Name).SetAuth(a)

	_, err = uc.Save(ctx)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return u.mailer.Send(p.Email, "Hello")
}
