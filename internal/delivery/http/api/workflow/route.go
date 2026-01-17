package workflow

import (
	"net/http"
	"workflow_engine/internal/delivery/http/middleware"
)

func RegisterRoutes(mux *http.ServeMux, handler *WorkFlowHandler, authMiddleware *middleware.AuthMiddleware) {
	mux.Handle("GET /api/document/{id}/signal", authMiddleware.Authenticate(http.HandlerFunc(handler.Signal)))
}
