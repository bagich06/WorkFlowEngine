package usecase

import (
	"context"
	"workflow_engine/internal/domain/entities"
)

type DocumentRepository interface {
	Create(ctx context.Context, document *entities.Document) (*entities.Document, error)
	GetByID(ctx context.Context, id int64) (*entities.Document, error)
	UpdateStatus(ctx context.Context, newStatus entities.DocumentStatus, id int64) error
	GetStatusByID(ctx context.Context, id int64) (entities.DocumentStatus, error)
}

type DocumentUseCase struct {
	docRepo DocumentRepository
}

func NewDocumentUseCase(docRepo DocumentRepository) *DocumentUseCase {
	return &DocumentUseCase{docRepo: docRepo}
}

func (uc *DocumentUseCase) Create(ctx context.Context, document *entities.Document) (*entities.Document, error) {
	created, err := uc.docRepo.Create(ctx, document)
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
