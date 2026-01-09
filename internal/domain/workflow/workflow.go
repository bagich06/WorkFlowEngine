package workflow

import (
	"workflow_engine/internal/domain/entities"
)

type StateMachine struct{}

func NewStateMachine() *StateMachine {
	return &StateMachine{}
}

func (sm *StateMachine) ValidateApproval(currentStatus entities.DocumentStatus, role entities.UserRole) error {
	if IsTerminalStatus(currentStatus) {
		return entities.ErrWorkflowAlreadyCompleted
	}

	expectedRole, ok := GetExpectedRole(currentStatus)
	if !ok {
		return entities.ErrInvalidTransition
	}

	if expectedRole != role {
		return entities.ErrNotYourTurn
	}

	return nil
}

func (sm *StateMachine) ProcessApproval(currentStatus entities.DocumentStatus, role entities.UserRole) (entities.DocumentStatus, error) {
	if err := sm.ValidateApproval(currentStatus, role); err != nil {
		return "", err
	}

	newStatus, ok := GetNextStatus(currentStatus, role)
	if !ok {
		return "", entities.ErrInvalidTransition
	}

	return newStatus, nil
}

func (sm *StateMachine) ProcessRejection(currentStatus entities.DocumentStatus, role entities.UserRole) (entities.DocumentStatus, error) {
	if err := sm.ValidateApproval(currentStatus, role); err != nil {
		return "", err
	}

	newStatus, ok := GetFailedStatus(currentStatus, role)
	if !ok {
		return "", entities.ErrInvalidTransition
	}

	return newStatus, nil
}
