package usecase

import (
	"context"
	"fmt"
	"mygo/http/oapi"
)

// CreateTask タスク作成
// タスクのDB登録と、タスクassign通知
func (u *UseCase) CreateTask(ctx context.Context, p *oapi.Task) error {
	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	user, err := tx.User.Get(ctx, 10)
	if err != nil {
		return err
	}

	a, err := user.QueryAuth().Only(ctx)
	if err != nil {
		return err
	}

	t, err := u.Repository.CreateTask(ctx, tx, p.Title, user, user)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	
	return u.Mailer.Send(a.Email, fmt.Sprintf("%s宛にID%dの「%s」を通知しました", a.Email, t.ID, t.Title))
}
