package workflow

import (
	"net/http"
	"workflow_engine/internal/delivery/http/middleware"
)

func RegisterRoutes(mux *http.ServeMux, handler *WorkflowHandler, authMiddleware *middleware.AuthMiddleware) {
	mux.Handle("POST /workflow/{id}/approve", authMiddleware.Authenticate(http.HandlerFunc(handler.Approve)))
	mux.Handle("POST /workflow/{id}/reject", authMiddleware.Authenticate(http.HandlerFunc(handler.Reject)))
	mux.Handle("GET /workflow/{id}/status", authMiddleware.Authenticate(http.HandlerFunc(handler.GetStatus)))
}
