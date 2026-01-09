package document

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"workflow_engine/internal/domain/entities"
	"workflow_engine/internal/domain/interfaces"
	"workflow_engine/pkg/validator"
)

type DocumentHandler struct {
	documentUseCase interfaces.DocumentUseCase
	validator       *validator.Validator
}

func NewDocumentHandler(documentUseCase interfaces.DocumentUseCase) *DocumentHandler {
	return &DocumentHandler{
		documentUseCase: documentUseCase,
	}
}

func (h *DocumentHandler) CreateDocument(w http.ResponseWriter, r *http.Request) {
	var req DocumentCreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.validator.Validate(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Status != entities.DocumentStatusStarted {
		http.Error(w, "Invalid initial status", http.StatusBadRequest)
		return
	}

	document := &entities.Document{
		Topic:  req.Topic,
		Status: req.Status,
	}

	documentID, err := h.documentUseCase.Create(r.Context(), document)
	if err != nil {
		if errors.Is(err, entities.ErrNotWorker) {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(documentID)
}

func (h *DocumentHandler) GetDocumentByID(w http.ResponseWriter, r *http.Request) {
	docIDStr := r.PathValue("id")
	docID, err := strconv.ParseInt(docIDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	document, err := h.documentUseCase.GetByID(r.Context(), docID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	docResp := DocumentResponse{
		ID:     document.ID,
		Topic:  document.Topic,
		Status: document.Status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docResp)
}
