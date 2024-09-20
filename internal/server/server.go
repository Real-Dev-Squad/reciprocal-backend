package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Real-Dev-Squad/reciprocal-backend/internal/config"
	"github.com/uptrace/bun"
)

type Server struct {
	db *bun.DB
}

func NewServer(ctx context.Context, db *bun.DB) *http.Server {
	NewServer := &Server{db}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      NewServer.RegisterRoutes(ctx),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
