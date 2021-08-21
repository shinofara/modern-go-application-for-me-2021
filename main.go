package main

import (
	"context"
	"contrib.go.opencensus.io/integrations/ocsql"
	"database/sql/driver"
	"entgo.io/ent/dialect"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"mygo/ent"
	"mygo/http/handler"
	oapi "mygo/http/openapi"
	entsql "entgo.io/ent/dialect/sql"
	"net/http"
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

func dsn() string {
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

	return mc.FormatDSN()
}

func main() {
	r := chi.NewRouter()
	mux := &handler.Handler{
		DB: Open(dsn()),
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
