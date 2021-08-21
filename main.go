package main

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"
	"log"
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
	db := database.Open(dsn())
	defer db.Close()

	container := dig.New()
	err := container.Provide(interfaces.NewDummyMailer)
	log.Println(err)
	err = container.Provide(handler.NewHandler)
	log.Println(err)
	err = container.Provide(database.Open)
	log.Println(err)
	err = container.Provide(dsn)
	log.Println(err)
	err = container.Invoke(Server)
	log.Println(err)
}

func Server(mux handler.Handler) error {
	r := chi.NewRouter()
	oapi.HandlerFromMux(&mux, r)
	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	log.Println("run: " + s.Addr)

	return s.ListenAndServe()
}
