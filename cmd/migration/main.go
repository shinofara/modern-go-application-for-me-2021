package main

import (
	"context"
	"github.com/go-sql-driver/mysql"
	"mygo/ent"
	"mygo/ent/migrate"
	"log"
)

func main() {
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
	ctx := context.Background()
	if err := client.Schema.Create(ctx, migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
	log.Print("ent sample done.")
}