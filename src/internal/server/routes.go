package server

import (
	"net/http"

	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/health"
)

func (s *Server) RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", health.HealthCheck)
	return mux
}
