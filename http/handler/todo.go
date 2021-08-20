package handler

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"mygo/ent"
	"net/http"
)

type Handler struct {
}

func (h Handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
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

	client, err := ent.Open("mysql", mc.FormatDSN(), entOptions...)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	todos, err := client.Todo.Query().All(ctx)
	if err != nil {
		log.Println(err.Error())
		fmt.Fprint(w, err)
	}

	fmt.Fprintf(w, "%v", todos)
	return
}

func (h Handler) Todo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
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

	client, err := ent.Open("mysql", mc.FormatDSN(), entOptions...)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()


	t := client.Todo.Create()
	t.SetTitle("hoge").SetDescription("hogehoge")
	if _, err := t.Save(ctx); err != nil {
		log.Println(err.Error())
		fmt.Fprint(w, err)
		return
	}


	fmt.Fprint(w, "add todo")
	return
}