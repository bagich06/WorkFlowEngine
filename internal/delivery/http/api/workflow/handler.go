package workflow

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"workflow_engine/internal/domain/entities/workflow"
)

type WorkFlowUseCase interface {
	HandleSignal(ctx context.Context, workflowID int64, signal workflow.WorkflowSignal) error
}

type WorkFlowHandler struct {
	uc WorkFlowUseCase
}

func NewWorkFlowHandler(uc WorkFlowUseCase) *WorkFlowHandler {
	return &WorkFlowHandler{uc: uc}
}

func (h *WorkFlowHandler) Signal(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	workflowID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid workflow id", http.StatusBadRequest)
		return
	}

	var req SignalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	err = h.uc.HandleSignal(
		r.Context(),
		workflowID,
		workflow.WorkflowSignal{
			Action: workflow.WorkflowAction(req.Action),
			Role:   string(req.Role),
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
