package document

import (
	"net/http"
	"workflow_engine/internal/delivery/http/middleware"
	"workflow_engine/internal/domain/entities"
)

func RegisterRoutes(mux *http.ServeMux, handler *DocumentHandler, authMiddleware *middleware.AuthMiddleware) {
	workerOnly := func(h http.HandlerFunc) http.Handler {
		return authMiddleware.Authenticate(
			authMiddleware.RequireRole(entities.RoleWorker)(http.HandlerFunc(h)),
		)
	}
	mux.Handle("POST /api/document", workerOnly(handler.CreateDocument))
	mux.Handle("GET /api/document/{id}", authMiddleware.Authenticate(http.HandlerFunc(handler.GetDocumentByID)))
	mux.Handle("PUT /api/document/status/{id}", authMiddleware.Authenticate(http.HandlerFunc(handler.UpdateDocumentStatus)))
}
