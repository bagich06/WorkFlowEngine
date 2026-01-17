package workflow

type WorkflowStatus string

const (
	WorkflowStatusRunning   WorkflowStatus = "running"
	WorkflowStatusCompleted WorkflowStatus = "completed"
	WorkflowStatusRejected  WorkflowStatus = "rejected"
)

type Workflow struct {
	ID       int64
	EntityID int64
	Step     int
	Status   WorkflowStatus
}
