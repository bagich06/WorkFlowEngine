package usecase

import (
	"context"
	"workflow_engine/internal/domain/entities"
	"workflow_engine/internal/domain/interfaces"
)

type DocumentUseCase struct {
	docRepo interfaces.DocumentRepository
}

func NewDocumentUseCase(docRepo interfaces.DocumentRepository) *DocumentUseCase {
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
