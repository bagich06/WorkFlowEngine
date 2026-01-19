package workflow

import "workflow_engine/internal/domain/entities"

type WorkflowStatus string
type WorkflowGroup string

const (
	WorkflowStatusRunning   WorkflowStatus = "running"
	WorkflowStatusCompleted WorkflowStatus = "completed"
	WorkflowStatusRejected  WorkflowStatus = "rejected"

	WorkflowGroupFirst WorkflowGroup = "first"
	WorkflowGroupFinal WorkflowGroup = "final"
)

type Workflow struct {
	ID           int64
	EntityID     int64
	Group        WorkflowGroup
	RolesStatus  map[entities.UserRole]bool
	GroupsStatus map[WorkflowGroup]bool
	Status       WorkflowStatus
}
