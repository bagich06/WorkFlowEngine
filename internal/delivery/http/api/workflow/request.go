package workflow

import "workflow_engine/internal/domain/entities/workflow"

type SignalRequest struct {
	Action workflow.WorkflowAction `json:"action"`
}
