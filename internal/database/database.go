package database

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/Real-Dev-Squad/reciprocal-backend/internal/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func New() *bun.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DbUsername, config.DbPassword, config.DbHost, strconv.Itoa(config.DbPort), config.Database)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	return bun.NewDB(sqldb, pgdialect.New())
}
