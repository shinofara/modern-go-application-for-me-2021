package repository

import (
	"context"

	"github.com/shinofara/example-go-2021/ent"
)

type Repository struct {
	DB *ent.Client
}

func NewRepository(db *ent.Client) *Repository {
	return &Repository{
		DB: db,
	}
}

// Create tasks作成時に関わるデータの永続化
func (r *Repository) CreateTask(ctx context.Context, tx *ent.Tx, title string, creator, assign *ent.User) (*ent.Task, error) {
	tc := tx.Task.Create()
	return tc.SetTitle(title).
		SetCreator(creator).
		SetAssign(assign).Save(ctx)
}
