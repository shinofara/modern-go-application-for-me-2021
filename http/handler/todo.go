package handler

import (
	"fmt"
	"log"
	"mygo/ent"
	"mygo/interfaces"
	"net/http"
)

type Handler struct {
	DB *ent.Client
	Mailer interfaces.MailerInterface
}

func NewHandler(db *ent.Client, mailer interfaces.MailerInterface) Handler {
	log.Println("NewHandler")
	return Handler{
		DB: db,
		Mailer: mailer,
	}
}

func (h Handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	todos, err := h.DB.Todo.Query().All(ctx)
	if err != nil {
		log.Println(err.Error())
		fmt.Fprint(w, err)
	}

	fmt.Fprintf(w, "%v", todos)
	return
}

func (h Handler) Todo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	t := h.DB.Todo.Create()
	t.SetTitle("hoge").SetDescription("hogehoge")
	if _, err := t.Save(ctx); err != nil {
		log.Println(err.Error())
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, "add todo")
	return
}