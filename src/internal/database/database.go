package database

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/config"
	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type healthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() healthResponse

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close()
}

type service struct {
	dbPool *pgxpool.Pool
}

var (
	dbInstance *service
)

func New() Service {
	if dbInstance != nil {
		return dbInstance
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DbUsername, config.DbPassword, config.DbHost, strconv.Itoa(config.DbPort), config.Database)

	dbPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		logger.Fatal("Error connecting to database: ", err)
	}

	dbInstance = &service{dbPool: dbPool}

	return dbInstance
}

func (s *service) Health() healthResponse {
	err := s.dbPool.Ping(context.Background())

	res := healthResponse{
		Status:  "up",
		Message: "",
	}

	if err != nil {
		logger.Error("Error pinging database: ", err)

		res.Status = "down"
		res.Message = "Unable to connect to database"

		return res
	}

	stats := s.dbPool.Stat()

	if stats.IdleConns() > 10 {
		res.Message = "Too many idle connections"
	}

	return res
}

func (s *service) Close() {
	s.dbPool.Close()
	logger.Info("Disconnected from database: ", config.Database)
}
