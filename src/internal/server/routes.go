package server

import (
	"context"
	"net/http"

	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/health"
)

func (s *Server) RegisterRoutes(ctx context.Context) *http.ServeMux {
	mux := http.NewServeMux()

	health.HealthRouteGroup(mux, ctx, s.db)

	return mux
}
