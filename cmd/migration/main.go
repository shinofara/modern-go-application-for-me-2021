package main

import (
	"context"
	"flag"
	"log"
	"mygo/config"
	"mygo/ent"
	"mygo/ent/migrate"
	"mygo/infrastructure/database"
	"os"
)

func main() {
	var dryrun bool
	var configPath string
	flag.BoolVar(&dryrun, "dryrun", true, "dryrun")
	flag.StringVar(&configPath, "config", "", "path to config yaml path")
	flag.Parse()

	entOptions := []ent.Option{}

	// 発行されるSQLをロギングするなら
	entOptions = append(entOptions, ent.Debug())

	cfg, err := config.New(configPath)
	if err != nil {
		panic(err)
	}

	client := database.NewClient(&cfg.DB)
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
