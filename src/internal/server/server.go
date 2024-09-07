package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	db *pgxpool.Pool
}

func NewServer(db *pgxpool.Pool) *http.Server {
	NewServer := &Server{db}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
