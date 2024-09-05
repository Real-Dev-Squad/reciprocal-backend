package middlewares

import (
	"fmt"
	"net/http"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Auth middleware")
		next(w, r)
	}
}

func Subscription(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Subscription middleware")
		next(w, r)
	}
}
