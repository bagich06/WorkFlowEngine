package entities

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user with this credentials already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidPassword   = errors.New("invalid password")

	ErrNotWorker               = errors.New("only workers are allowed to do this")
	ErrWorkflowAlreadyExists   = errors.New("workflow already exists")
	ErrRoleNotAllowed          = errors.New("role not allowed")
	ErrInvalidAction           = errors.New("invalid action")
	ErrWorkflowAlreadyFinished = errors.New("workflow already finished")
	ErrRoleAlreadyFinished     = errors.New("role already finished")
	//ErrNotManager   = errors.New("only managers are allowed to do this")
	//ErrNotEconomist = errors.New("only economists is allowed to do this")
	//ErrNotBoss      = errors.New("only boss is allowed to do this")
	//ErrFailedStatus = errors.New("process failed")
	//ErrNotReady     = errors.New("process not ready")
	//ErrDocumentNotFound         = errors.New("document not found")
	//ErrInvalidTransition        = errors.New("invalid status transition")
	//ErrNotYourTurn              = errors.New("it's not your turn to approve")
	//ErrWorkflowAlreadyCompleted = errors.New("workflow already completed")
)
