package database

import (
	"context"
	"contrib.go.opencensus.io/integrations/ocsql"
	"database/sql"
	"database/sql/driver"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-sql-driver/mysql"
	"mygo/ent"
)

type Config struct {
	User string
	Passwd string
	Net string
	Host string
	Port string
	DBName string `yaml:"db_name"`
}

type connector struct {
	dsn string
}

func (c connector) Connect(context.Context) (driver.Conn, error) {
	return c.Driver().Open(c.dsn)
}

func (connector) Driver() driver.Driver {
	return ocsql.Wrap(
		mysql.MySQLDriver{},
		ocsql.WithAllTraceOptions(),
		ocsql.WithRowsClose(false),
		ocsql.WithRowsNext(false),
		ocsql.WithDisableErrSkip(true),
	)
}

// Open new connection and start stats recorder.
func NewClient(cfg *Config) *ent.Client {
	mc := mysql.Config{
		User:                 cfg.User,
		Passwd:               cfg.Passwd,
		Net:                  cfg.Net,
		Addr:                 cfg.Host + ":" + cfg.Port,
		DBName:               cfg.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}


	db := sql.OpenDB(connector{mc.FormatDSN()})
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.MySQL, db)
	return ent.NewClient(ent.Driver(drv))
}