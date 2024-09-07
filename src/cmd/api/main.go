package main

import (
	"log"

	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/server"
)

func main() {
	server := server.NewServer()

	log.Println("Server running on port 4000")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
