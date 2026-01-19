package usecase

import (
	"context"
	"workflow_engine/internal/domain/entities"
	"workflow_engine/internal/domain/entities/workflow"
)

type DocumentRepository interface {
	Create(ctx context.Context, document *entities.Document) (*entities.Document, error)
	GetByID(ctx context.Context, id int64) (*entities.Document, error)
	UpdateStatus(ctx context.Context, newStatus entities.DocumentStatus, id int64) error
	GetStatusByID(ctx context.Context, id int64) (entities.DocumentStatus, error)
}

type DocumentUseCase struct {
	docRepo DocumentRepository
	wf      WorkflowRepository
}

func NewDocumentUseCase(docRepo DocumentRepository, wf WorkflowRepository) *DocumentUseCase {
	return &DocumentUseCase{docRepo: docRepo, wf: wf}
}

func (uc *DocumentUseCase) Create(ctx context.Context, document *entities.Document) (*entities.Document, error) {
	document.Status = entities.DocumentStatusStarted
	created, err := uc.docRepo.Create(ctx, document)
	if err != nil {
		return nil, err
	}

	wf := &workflow.Workflow{
		EntityID:     created.ID,
		Group:        workflow.WorkflowGroupFirst,
		RolesStatus:  nil,
		GroupsStatus: nil,
		Status:       workflow.WorkflowStatusRunning,
	}

	err = uc.wf.Create(ctx, wf)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (uc *DocumentUseCase) GetByID(ctx context.Context, id int64) (*entities.Document, error) {
	doc, err := uc.docRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
