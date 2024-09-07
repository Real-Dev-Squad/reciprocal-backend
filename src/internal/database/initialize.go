package database

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/config"
	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type service struct {
	dbPool *pgxpool.Pool
}

func Initialize(ctx context.Context) *pgxpool.Pool {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DbUsername, config.DbPassword, config.DbHost, strconv.Itoa(config.DbPort), config.Database)
	connectionPool, err := pgxpool.New(ctx, connStr)

	if err != nil {
		logger.Error("Error connecting to database: ", err)
	}

	return connectionPool
}
