package main

import (
	"fmt"
	"net/http"

	"github.com/Real-Dev-Squad/reciprocal-backend/src/handlers"
	"github.com/Real-Dev-Squad/reciprocal-backend/src/middlewares"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users/greet/{name}", middlewares.Auth(middlewares.Subscription(handlers.GreetUser)))
	mux.HandleFunc("POST /users", handlers.CreateUser)

	port := ":8080"
	fmt.Println("Server is running on port", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		panic(err)
	}
}
