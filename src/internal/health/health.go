package health

import (
	"encoding/json"
	"net/http"

	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/database"
	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/logger"
)

type health struct {
	db database.Service
}

func (h *health) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse, err := json.Marshal(h.db.Health())

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		logger.Error("Error marshalling health response: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonResponse))
}

func HealthRouteGroup(mux *http.ServeMux, db database.Service) {
	healthService := &health{db}
	mux.HandleFunc("GET /health", healthService.healthHandler)
}
