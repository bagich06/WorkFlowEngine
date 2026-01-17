package workflow

type WorkflowAction string

const (
	WorkflowActionApprove WorkflowAction = "approve"
	WorkflowActionReject  WorkflowAction = "reject"
)

type WorkflowSignal struct {
	Action WorkflowAction
	Role   string
}
