package mysql

import (
	"context"
	"log"
	"mygo/ent"
	"mygo/ent/migrate"
	"mygo/infrastructure/database"
	"testing"

	"github.com/DATA-DOG/go-txdb"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-sql-driver/mysql"
)

func init() {
	client := database.NewClient(&database.Config{
		Host:   "127.0.0.1",
		Net:    "tcp",
		Port:   "3306",
		DBName: "example",
		User:   "root",
	})
	defer client.Close()
	ctx := context.Background()

	if err := client.Schema.Create(
		ctx,
		migrate.WithForeignKeys(false),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Panicf("failed printing schema changes: %v", err)
	}
}

func NewTestClient(t *testing.T) *ent.Client {
	cfg := &database.Config{
		Host:   "127.0.0.1",
		Net:    "tcp",
		Port:   "3306",
		DBName: "example",
		User:   "root",
	}

	mc := mysql.Config{
		User:                 cfg.User,
		Passwd:               cfg.Passwd,
		Net:                  cfg.Net,
		Addr:                 cfg.Host + ":" + cfg.Port,
		DBName:               cfg.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	driverName := "mysql2"
	txdb.Register(driverName, "mysql", mc.FormatDSN())
	drv, err := entsql.Open(driverName, "identifier")
	if err != nil {
		t.Fatal(err)
	}

	o := []ent.Option{
		ent.Driver(drv),
	}
	return ent.NewClient(o...)
}
