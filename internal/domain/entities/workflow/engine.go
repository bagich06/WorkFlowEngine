package workflow

import "errors"

func ApplySignal(workflow *Workflow, definition WorkflowDefinition, signal WorkflowSignal) error {
	if workflow.Status != WorkflowStatusRunning {
		return errors.New("workflow already finished")
	}

	if workflow.Step >= len(definition.Steps) {
		return errors.New("invalid workflow step")
	}

	currentStep := definition.Steps[workflow.Step]

	if signal.Role != currentStep.AllowedRole {
		return errors.New("role not allowed for this step")
	}

	if signal.Action == WorkflowActionReject {
		workflow.Status = WorkflowStatusRejected
		return nil
	}

	workflow.Step++

	if workflow.Step >= len(definition.Steps) {
		workflow.Status = WorkflowStatusCompleted
	}

	return nil
}
