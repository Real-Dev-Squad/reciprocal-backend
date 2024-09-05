package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Real-Dev-Squad/reciprocal-backend/src/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("User created:", user.Name)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GreetUser(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	age := r.URL.Query().Get("age")

	if age != "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello, " + name + "! You are " + age + " years old.",
		})

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello, " + name + "!",
	})
}
