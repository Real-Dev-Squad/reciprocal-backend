package database

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/config"
	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/logger"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

type service struct {
	db *sql.DB
}

var (
	dbInstance *service
)

func New() Service {
	if dbInstance != nil {
		return dbInstance
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DbUsername, config.DbPassword, config.DbHost, strconv.Itoa(config.DbPort), config.Database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Fatal("Error connecting to database: ", err)
	}

	dbInstance = &service{db: db}

	return dbInstance
}

func (s *service) Health() map[string]string {
	return map[string]string{
		"status": "ok",
	}
}

/*
* Closes the database connection.
* @returns -  nil if successfully closed the connection, error if the connection cannot be closed.
 */
func (s *service) Close() error {
	logger.Info("Disconnected from database: ", config.Database)
	return s.db.Close()
}
