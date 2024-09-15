package health

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type healthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type health struct {
	db *pgxpool.Pool
}

func (h *health) healthHandler(w http.ResponseWriter, r *http.Request) {
	err := h.db.Ping(r.Context())

	res := healthResponse{
		Status:  "Up",
		Message: "All systems operational!",
	}

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		res.Status = "Down"
		res.Message = "Unable to connect to database"

		jsonRes, err := json.Marshal(res)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonRes)
		return
	}

	jsonRes, err := json.Marshal(res)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

func HealthRouteGroup(mux *http.ServeMux, db *pgxpool.Pool) {
	healthService := &health{db}
	mux.HandleFunc("GET /health", healthService.healthHandler)
}
