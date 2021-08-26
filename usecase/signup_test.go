package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/shinofara/example-go-2021/openapi"

	"github.com/shinofara/example-go-2021/ent/auth"
	mock_mailer "github.com/shinofara/example-go-2021/mock"
	"github.com/shinofara/example-go-2021/repository"
	"github.com/shinofara/example-go-2021/testsupport/mysql"

	"github.com/golang/mock/gomock"
)

func TestUseCase_Signup(t *testing.T) {
	client := mysql.NewTestClient(t)
	defer client.Close()
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	m := mock_mailer.NewMockMailerInterface(ctrl)
	m.EXPECT().Send("test1@example.com", "Hello").Return(nil)
	m.EXPECT().Send("test2@example.com", "Hello").Return(errors.New("send error"))
	repo := repository.NewRepository(client)

	uc := &UseCase{
		db:         client,
		mailer:     m,
		repository: repo,
	}

	t.Run("Success", func(t *testing.T) {
		if err := uc.Signup(ctx, &openapi.SignupRequest{
			Name:     "test",
			Email:    "test1@example.com",
			Password: "test",
		}); err != nil {
			t.Errorf("Signup() error = %v, want nil", err)
		}
	})

	t.Run("Failed send email", func(t *testing.T) {
		if err := uc.Signup(ctx, &openapi.SignupRequest{
			Name:     "test",
			Email:    "test2@example.com",
			Password: "test",
		}); err == nil {
			t.Errorf("Signup() error = %v, want err", err)
		}

		if exists := client.Auth.Query().Where(auth.Email("test2@example.com")).ExistX(ctx); !exists {
			t.Errorf("Auth not exists, want true")
		}
	})

	t.Run("Failed because the user already exists", func(t *testing.T) {
		a := client.Auth.Create().SetEmail("test3@example.com").SetPassword("test3").SaveX(ctx)
		u := client.User.Create().SetName("test3").SetAuth(a).SaveX(ctx)

		if err := uc.Signup(ctx, &openapi.SignupRequest{
			Name:     u.Name,
			Email:    a.Email,
			Password: a.Password,
		}); err == nil {
			t.Errorf("Signup() error = %v, want err", err)
		}

		if cnt := client.Auth.Query().Where(auth.Email(a.Email)).CountX(ctx); cnt != 1 {
			t.Errorf("Auth cnt %d, want 1", cnt)
		}
	})
}
