package usecase

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/shinofara/modern-go-application-for-me-2021/openapi"
)

// CreateTask タスク作成
// タスクのDB登録と、タスクassign通知
func (u *UseCase) CreateTask(ctx context.Context, p *openapi.Task) error {
	// 一旦意味も無くtransactionで書いてる
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

	user, err := tx.User.Get(ctx, 16)
	if err != nil {
		return err
	}

	a, err := user.QueryAuth().Only(ctx)
	if err != nil {
		return err
	}

	t, err := u.repository.CreateTask(ctx, tx, p.Title, user, user)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return u.mailer.Send(a.Email, fmt.Sprintf("%s宛にID%dの「%s」を通知しました", a.Email, t.ID, t.Title))
}
