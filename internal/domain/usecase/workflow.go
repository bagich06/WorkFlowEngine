package usecase

import (
	"context"
	"workflow_engine/internal/domain/entities"
	"workflow_engine/internal/domain/entities/workflow"
)

type WorkflowRepository interface {
	GetByID(ctx context.Context, id int64) (*workflow.Workflow, error)
	Save(ctx context.Context, wf *workflow.Workflow) error
}

type WorkFlowUseCase struct {
	workflowRepo WorkflowRepository
	documentRepo DocumentRepository
}

func NewWorkFlowRepository(workflowRepo WorkflowRepository, documentRepo DocumentRepository) *WorkFlowUseCase {
	return &WorkFlowUseCase{
		workflowRepo: workflowRepo,
		documentRepo: documentRepo,
	}
}

func (uc *WorkFlowUseCase) HandleSignal(ctx context.Context, workflowID int64, signal workflow.WorkflowSignal) error {

	createdWorkflow, err := uc.workflowRepo.GetByID(ctx, workflowID)
	if err != nil {
		return err
	}

	err = workflow.ApplySignal(
		createdWorkflow,
		workflow.DocumentApprovalWorkflow,
		signal,
	)
	if err != nil {
		return err
	}

	switch createdWorkflow.Status {

	case workflow.WorkflowStatusCompleted:
		if err := uc.documentRepo.UpdateStatus(
			ctx,
			entities.DocumentStatusApproved,
			createdWorkflow.EntityID,
		); err != nil {
			return err
		}

	case workflow.WorkflowStatusRejected:
		if err := uc.documentRepo.UpdateStatus(
			ctx,
			entities.DocumentStatusRejected,
			createdWorkflow.EntityID,
		); err != nil {
			return err
		}
	}

	return uc.workflowRepo.Save(ctx, createdWorkflow)
}
