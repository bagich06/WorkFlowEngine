package workflow

import (
	"errors"
	"workflow_engine/internal/domain/entities"
)

func ApplySignal(workflow *Workflow, definition WorkflowDefinition, signal WorkflowSignal) error {
	if workflow.Status != WorkflowStatusRunning {
		return errors.New("workflow already finished")
	}

	currentGroup := workflow.Group

	var groupSteps []WorkflowStep
	roleAllowed := false
	for _, step := range definition.Steps {
		if step.ParallelGroup == currentGroup {
			groupSteps = append(groupSteps, step)
			if step.AllowedRole == signal.Role {
				roleAllowed = true
			}
		}
	}

	if !roleAllowed {
		return entities.ErrRoleNotAllowed
	}

	if signal.Action == WorkflowActionReject {
		workflow.Status = WorkflowStatusRejected
		if workflow.RolesStatus == nil {
			workflow.RolesStatus = make(map[entities.UserRole]bool)
		}
		workflow.RolesStatus[signal.Role] = false
		return nil
	}

	if signal.Action != WorkflowActionApprove {
		return entities.ErrInvalidAction
	}

	if workflow.RolesStatus == nil {
		workflow.RolesStatus = make(map[entities.UserRole]bool)
	}
	workflow.RolesStatus[signal.Role] = true

	allDone := true
	for _, step := range groupSteps {
		if !workflow.RolesStatus[step.AllowedRole] {
			allDone = false
			break
		}
	}

	if allDone {
		if workflow.GroupsStatus == nil {
			workflow.GroupsStatus = make(map[WorkflowGroup]bool)
		}
		workflow.GroupsStatus[currentGroup] = true
	}

	var foundCurrent bool
	var nextGroup WorkflowGroup
	for _, step := range definition.Steps {
		if step.ParallelGroup == currentGroup {
			foundCurrent = true
			continue
		}
		if foundCurrent && !workflow.GroupsStatus[step.ParallelGroup] {
			nextGroup = step.ParallelGroup
			break
		}
	}

	if allDone {
		if nextGroup != "" {
			workflow.Group = nextGroup
			workflow.RolesStatus = make(map[entities.UserRole]bool)
		} else {
			workflow.Status = WorkflowStatusCompleted
		}
	}

	return nil
}
