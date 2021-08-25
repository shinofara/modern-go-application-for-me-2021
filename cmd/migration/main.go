package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/shinofara/example-go-2021/config"
	"github.com/shinofara/example-go-2021/ent"
	"github.com/shinofara/example-go-2021/ent/migrate"
	"github.com/shinofara/example-go-2021/infrastructure/database"
)

func main() {
	var dryrun bool
	var configPath string
	flag.BoolVar(&dryrun, "dryrun", true, "dryrun")
	flag.StringVar(&configPath, "config", "", "path to config yaml path")
	flag.Parse()

	entOptions := []ent.Option{
		ent.Debug(),
	}

	cfg, err := config.New(configPath)
	if err != nil {
		panic(err)
	}

	client := database.NewClient(cfg.DB, entOptions...)
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
