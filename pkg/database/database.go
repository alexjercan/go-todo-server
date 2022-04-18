package database

import (
	"crypto/tls"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var (
	database bun.DB
)

type Connection struct {
	Address  string
	User     string
	Password string
	Database string
}

func Connect(connection Connection) {
	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(connection.Address),
		pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		pgdriver.WithUser(connection.User),
		pgdriver.WithPassword(connection.Password),
		pgdriver.WithDatabase(connection.Database),
		pgdriver.WithApplicationName("myapp"),
		pgdriver.WithTimeout(5*time.Second),
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(5*time.Second),
		pgdriver.WithWriteTimeout(5*time.Second),
	)

	sqldb := sql.OpenDB(pgconn)

	database = *bun.NewDB(sqldb, pgdialect.New())
}
