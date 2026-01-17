package workflow

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"workflow_engine/internal/delivery/http/middleware"
	"workflow_engine/internal/domain/entities"
)

type WorkflowUseCase interface {
	Approve(ctx context.Context, docID int64, userID int64) error
	Reject(ctx context.Context, docID int64, userID int64) error
	GetDocumentStatus(ctx context.Context, docID int64) (entities.DocumentStatus, error)
}

type WorkflowHandler struct {
	workflowUseCase WorkflowUseCase
}

func NewWorkflowHandler(workflowUseCase WorkflowUseCase) *WorkflowHandler {
	return &WorkflowHandler{
		workflowUseCase: workflowUseCase,
	}
}

func (h *WorkflowHandler) Approve(w http.ResponseWriter, r *http.Request) {
	docIDStr := r.PathValue("id")
	docID, err := strconv.ParseInt(docIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid document id", http.StatusBadRequest)
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "user not found in context", http.StatusUnauthorized)
		return
	}

	err = h.workflowUseCase.Approve(r.Context(), docID, userID)
	if err != nil {
		statusCode := h.mapErrorToStatusCode(err)
		http.Error(w, err.Error(), statusCode)
		return
	}

	status, _ := h.workflowUseCase.GetDocumentStatus(r.Context(), docID)

	resp := WorkflowResponse{
		Message: "Document approved successfully",
		Status:  status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *WorkflowHandler) Reject(w http.ResponseWriter, r *http.Request) {
	docIDStr := r.PathValue("id")
	docID, err := strconv.ParseInt(docIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid document id", http.StatusBadRequest)
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "user not found in context", http.StatusUnauthorized)
		return
	}

	err = h.workflowUseCase.Reject(r.Context(), docID, userID)
	if err != nil {
		statusCode := h.mapErrorToStatusCode(err)
		http.Error(w, err.Error(), statusCode)
		return
	}

	status, _ := h.workflowUseCase.GetDocumentStatus(r.Context(), docID)

	resp := WorkflowResponse{
		Message: "Document rejected",
		Status:  status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *WorkflowHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	docIDStr := r.PathValue("id")
	docID, err := strconv.ParseInt(docIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid document id", http.StatusBadRequest)
		return
	}

	status, err := h.workflowUseCase.GetDocumentStatus(r.Context(), docID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := StatusResponse{
		DocumentID: docID,
		Status:     status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *WorkflowHandler) mapErrorToStatusCode(err error) int {
	switch {
	case errors.Is(err, entities.ErrDocumentNotFound):
		return http.StatusNotFound
	case errors.Is(err, entities.ErrUserNotFound):
		return http.StatusUnauthorized
	case errors.Is(err, entities.ErrNotYourTurn):
		return http.StatusForbidden
	case errors.Is(err, entities.ErrWorkflowAlreadyCompleted):
		return http.StatusConflict
	case errors.Is(err, entities.ErrInvalidTransition):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
