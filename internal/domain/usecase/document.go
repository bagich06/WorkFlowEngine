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
	document.Status = entities.DocumentStatusStarted
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

func (uc *DocumentUseCase) UpdateStatus(ctx context.Context, newStatus entities.DocumentStatus, docID int64, userRole entities.UserRole) error {
	if (newStatus == entities.DocumentStatusManagerConfirmed || newStatus == entities.DocumentStatusManagerFailed) && userRole != entities.RoleManager {
		return entities.ErrNotManager
	}

	if (newStatus == entities.DocumentStatusEconomistConfirmed || newStatus == entities.DocumentStatusEconomistFailed) && (userRole != entities.RoleEconomist) {
		return entities.ErrNotEconomist
	}

	if (newStatus == entities.DocumentStatusBossConfirmed || newStatus == entities.DocumentStatusBossFailed) && (userRole != entities.RoleBoss) {
		return entities.ErrNotBoss
	}

	if newStatus == entities.DocumentStatusManagerFailed || newStatus == entities.DocumentStatusEconomistFailed || newStatus == entities.DocumentStatusBossFailed {
		return entities.ErrFailedStatus
	}

	err := uc.docRepo.UpdateStatus(ctx, newStatus, docID)
	if err != nil {
		return err
	}

	return nil
}
