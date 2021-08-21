package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"mygo/ent"
	"mygo/http/handler"
	oapi "mygo/http/openapi"
	"net/http"
)

func connect() (*ent.Client, error){
	entOptions := []ent.Option{}
	// 発行されるSQLをロギングするなら
	entOptions = append(entOptions, ent.Debug())
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

	return ent.Open("mysql", mc.FormatDSN(), entOptions...)
}

func main() {
	r := chi.NewRouter()
	db, err := connect()
	if err != nil {
		panic(err)
	}
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
