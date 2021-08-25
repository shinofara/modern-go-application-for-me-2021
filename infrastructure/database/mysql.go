package database

import (
	"mygo/ent"

	"github.com/XSAM/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-sql-driver/mysql"
)

type Config struct {
	User   string
	Passwd string
	Net    string
	Host   string
	Port   string
	DBName string `yaml:"db_name"`
}

// Open new connection and start stats recorder.
func NewClient(cfg *Config, opts ...ent.Option) *ent.Client {
	mc := mysql.Config{
		User:                 cfg.User,
		Passwd:               cfg.Passwd,
		Net:                  cfg.Net,
		Addr:                 cfg.Host + ":" + cfg.Port,
		DBName:               cfg.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	driverName, err := otelsql.Register("mysql", semconv.DBSystemMySQL.Value.AsString())
	if err != nil {
		panic(err)
	}

	drv, err := entsql.Open(driverName, mc.FormatDSN())
	if err != nil {
		panic(err)
	}

	o := []ent.Option{
		ent.Driver(drv),
	}

	o = append(o, opts...)
	return ent.NewClient(o...)
}
