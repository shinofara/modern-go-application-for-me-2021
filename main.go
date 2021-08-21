package main

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"
	"log"
	"mygo/ent"
	"mygo/http/handler"
	oapi "mygo/http/openapi"
	"github.com/go-sql-driver/mysql"
	"mygo/infrastructure/database"
	"mygo/interfaces"

	"net/http"
)


func dsn() string {
	// サンプルなのでここにハードコーディングしてます。
	mc := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "db" + ":" + "3306",
		DBName:               "example",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	return mc.FormatDSN()
}

func main() {
	provides := []interface{}{
		interfaces.NewDummyMailer,
		handler.NewHandler,
		database.Open,
		dsn,
	}

	container := dig.New()
	for _, p := range provides {
		if err := container.Provide(p); err != nil {
			panic(err)
		}
	}

	if err := container.Invoke(Server); err != nil {
		panic(err)
	}
}

func Server(mux handler.Handler, db *ent.Client) error {
	defer func() {
		db.Close()
		log.Println("DB Close")
	}()

	r := chi.NewRouter()
	oapi.HandlerFromMux(&mux, r)
	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	log.Println("run: " + s.Addr)

	return s.ListenAndServe()
}
