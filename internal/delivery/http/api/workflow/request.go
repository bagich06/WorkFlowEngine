package workflow

import "workflow_engine/internal/domain/entities"

type SignalRequest struct {
	Action string            `json:"action"`
	Role   entities.UserRole `json:"role"`
}
