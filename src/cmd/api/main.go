package main

import (
	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/logger"
	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/server"
)

func main() {
	server := server.NewServer()

	logger.Info("Server running on port 4000")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
