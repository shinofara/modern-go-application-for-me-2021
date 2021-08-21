package main

import (
	"github.com/go-chi/chi/v5"
	"mygo/http/handler"
	oapi "mygo/http/openapi"
	"github.com/go-sql-driver/mysql"
	"mygo/infrastructure/database"

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
	r := chi.NewRouter()

	db := database.Open(dsn())
	defer db.Close()

	mux := &handler.Handler{
		DB: db,
	}
	oapi.HandlerFromMux(mux, r)
	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
