package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/alexjercan/go-todo-server/pkg/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var (
	database bun.DB
)

type Connection struct {
	Address  string
	Port     string
	User     string
	Password string
	Database string
}

func Connect(connection Connection) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", connection.User, connection.Password, connection.Database, connection.Port, connection.Address)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	database = *bun.NewDB(sqldb, pgdialect.New())
}

func GetItems(ctx context.Context) ([]models.Item, error) {
	items := []models.Item{}

	err := database.NewSelect().Model(&items).Scan(ctx)

	return items, err
}

func InsertItem(ctx context.Context, item *models.Item) (*models.Item, error) {
	_, err := database.NewInsert().Model(item).Exec(ctx)

	return item, err
}

func UpdateItem(ctx context.Context, item *models.Item) (*models.Item, error) {
	_, err := database.NewUpdate().Model(item).Exec(ctx)

	return item, err
}

func DeleteItem(ctx context.Context, item *models.Item) (*models.Item, error) {
	_, err := database.NewDelete().Model(item).Exec(ctx)

	return item, err
}
