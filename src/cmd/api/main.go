package main

import (
	"context"

	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/config"
	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/database"
	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/logger"
	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/server"
)

func main() {
	// Create a context that we can cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create a database connection
	db := database.New()
	defer db.Close()

	// Create a new server
	server := server.NewServer(ctx, db)
	logger.Info("Server running on port", config.Port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
