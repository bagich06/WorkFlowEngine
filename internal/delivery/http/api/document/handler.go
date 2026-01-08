package document

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"workflow_engine/internal/delivery/http/middleware"
	"workflow_engine/internal/domain/entities"
	"workflow_engine/internal/domain/interfaces"
)

type DocumentHandler struct {
	documentUseCase interfaces.DocumentUseCase
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		Topic:  document.Topic,
		Status: document.Status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docResp)
}

func (h *DocumentHandler) UpdateDocumentStatus(w http.ResponseWriter, r *http.Request) {
	userRole, ok := middleware.GetRoleFromContext(r.Context())
	if !ok {
		http.Error(w, "Role not found in context", http.StatusUnauthorized)
		return
	}

	docIDStr := r.PathValue("id")
	docID, err := strconv.ParseInt(docIDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req DocumentUpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.documentUseCase.UpdateStatus(r.Context(), req.Status, docID, userRole)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := UpdateDocumentResponse{
		RespStr: "Status updated successfully",
		Status:  req.Status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
