package main

import (
	"context"
	"flag"
	"github.com/go-sql-driver/mysql"
	"mygo/ent"
	"mygo/ent/migrate"
	"log"
	"os"
)

func main() {
	var dryrun bool
	flag.BoolVar(&dryrun, "dryrun", true, "dryrun")
	flag.Parse()

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

	client, err := ent.Open("database", mc.FormatDSN(), entOptions...)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	if dryrun {
		if err := client.Schema.WriteTo(ctx, os.Stdout, migrate.WithForeignKeys(false)); err != nil {
			log.Fatalf("failed printing schema changes: %v", err)
		}
		os.Exit(0)
	}

	if err := client.Schema.Create(
		ctx,
		migrate.WithForeignKeys(false),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
	log.Print("ent sample done.")
}