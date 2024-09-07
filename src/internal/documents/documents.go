package documents

import (
	"encoding/json"
	"net/http"

	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/database"
	"github.com/Real-Dev-Squad/reciprocal-backend/src/internal/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type documents struct {
	db *pgxpool.Pool
}

func DocumentsRouteGroup(mux *http.ServeMux, db *pgxpool.Pool) {
	documentsService := &documents{db}
	mux.HandleFunc("GET /documents", documentsService.documentsHandler)
}

func (d *documents) documentsHandler(w http.ResponseWriter, r *http.Request) {
	queries := database.New(d.db)
	documents, err := queries.GetAllDocuments(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error("Error getting documents: ", err)
		return
	}

	jsonRes, err := json.Marshal(toDocumentsResponse(documents))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error("Error marshalling documents: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

func toDocumentsResponse(documents []database.Document) GetAllDocumentsResponse {
	return GetAllDocumentsResponse{Documents: documents}
}
