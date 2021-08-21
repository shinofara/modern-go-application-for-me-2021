package database

import (
	"context"
	"contrib.go.opencensus.io/integrations/ocsql"
	"database/sql/driver"
	"entgo.io/ent/dialect"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"mygo/ent"
	entsql "entgo.io/ent/dialect/sql"
)

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
func Open(dsn string) *ent.Client {
	db := sql.OpenDB(connector{dsn})
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.MySQL, db)
	return ent.NewClient(ent.Driver(drv))
}