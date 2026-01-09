package workflow

import (
	"context"
	"workflow_engine/internal/domain/entities"
	"workflow_engine/internal/domain/interfaces"
	"workflow_engine/internal/domain/workflow"
)

type WorkflowUseCase struct {
	docRepo      interfaces.DocumentRepository
	userRepo     interfaces.UserRepository
	stateMachine *workflow.StateMachine
}

func NewWorkflowUseCase(
	docRepo interfaces.DocumentRepository,
	userRepo interfaces.UserRepository,
) *WorkflowUseCase {
	return &WorkflowUseCase{
		docRepo:      docRepo,
		userRepo:     userRepo,
		stateMachine: workflow.NewStateMachine(),
	}
}

func (uc *WorkflowUseCase) Approve(ctx context.Context, docID int64, userID int64) error {
	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return entities.ErrUserNotFound
	}

	currentStatus, err := uc.docRepo.GetStatusByID(ctx, docID)
	if err != nil {
		return entities.ErrDocumentNotFound
	}

	newStatus, err := uc.stateMachine.ProcessApproval(currentStatus, user.Role)
	if err != nil {
		return err
	}

	if err := uc.docRepo.UpdateStatus(ctx, newStatus, docID); err != nil {
		return err
	}

	if newStatus == entities.DocumentStatusBossConfirmed {
		return uc.docRepo.UpdateStatus(ctx, entities.DocumentStatusSuccess, docID)
	}

	return nil
}

func (uc *WorkflowUseCase) Reject(ctx context.Context, docID int64, userID int64) error {
	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return entities.ErrUserNotFound
	}

	currentStatus, err := uc.docRepo.GetStatusByID(ctx, docID)
	if err != nil {
		return entities.ErrDocumentNotFound
	}

	newStatus, err := uc.stateMachine.ProcessRejection(currentStatus, user.Role)
	if err != nil {
		return err
	}

	return uc.docRepo.UpdateStatus(ctx, newStatus, docID)
}

func (uc *WorkflowUseCase) GetDocumentStatus(ctx context.Context, docID int64) (entities.DocumentStatus, error) {
	return uc.docRepo.GetStatusByID(ctx, docID)
}
