package handler

import (
	"net/http"

	"github.com/shinofara/modern-go-application-for-me-2021/ent"
	"github.com/shinofara/modern-go-application-for-me-2021/infrastructure/mailer"
	"github.com/shinofara/modern-go-application-for-me-2021/repository"
	"github.com/shinofara/modern-go-application-for-me-2021/usecase"

	"github.com/goccy/go-json"
	"go.uber.org/dig"
)

type Handler struct {
	db         *ent.Client
	mailer     mailer.MailerInterface
	useCase    *usecase.UseCase
	repository *repository.Repository
}

func NewHandler(p struct {
	dig.In

	DB         *ent.Client
	Mailer     mailer.MailerInterface
	UseCase    *usecase.UseCase
	Repository *repository.Repository
}) Handler {
	return Handler{
		db:         p.DB,
		mailer:     p.Mailer,
		useCase:    p.UseCase,
		repository: p.Repository,
	}
}

type response struct {
	status int
}

func (Handler) response() *response {
	return &response{
		status: http.StatusOK,
	}
}

func (r response) setStatus(s int) *response {
	rr := r
	rr.status = s
	return &rr
}

func (r response) json(w http.ResponseWriter, a interface{}) {
	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}
