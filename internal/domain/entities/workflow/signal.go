package workflow

import "workflow_engine/internal/domain/entities"

type WorkflowAction string

const (
	WorkflowActionApprove WorkflowAction = "approve"
	WorkflowActionReject  WorkflowAction = "reject"
)

type WorkflowSignal struct {
	Action WorkflowAction
	Role   entities.UserRole
}
