package server

import (
	"net/http"

	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/documents"
	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/health"
)

func (s *Server) RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	health.HealthRouteGroup(mux, s.db)
	documents.DocumentsRouteGroup(mux, s.db)

	return mux
}
