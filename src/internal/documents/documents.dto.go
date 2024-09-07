package documents

import "github.com/Real-Dev-Squad/reciprocal-backend/src/internal/database"

type GetAllDocumentsResponse struct {
	Documents []database.Document `json:"documents"`
}
