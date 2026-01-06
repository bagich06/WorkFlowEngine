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

func (uc *DocumentUseCase) Create(ctx context.Context, document *entities.Document, user *entities.User) (*entities.Document, error) {
	if user.Role != entities.RoleWorker {
		return nil, entities.ErrNotWorker
	}

	document.Status = entities.DocumentStatusStarted
	created, err := uc.docRepo.Create(ctx, document)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (uc *DocumentUseCase) UpdateStatus(ctx context.Context, newStatus entities.DocumentStatus, user entities.User) error {
	if (newStatus == entities.DocumentStatusManagerConfirmed || newStatus == entities.DocumentStatusManagerFailed) && user.Role != entities.RoleManager {
		return entities.ErrNotManager
	}

	if (newStatus == entities.DocumentStatusEconomistConfirmed || newStatus == entities.DocumentStatusEconomistFailed) && (user.Role != entities.RoleEconomist) {
		return entities.ErrNotEconomist
	}

	if (newStatus == entities.DocumentStatusBossConfirmed || newStatus == entities.DocumentStatusBossFailed) && (user.Role != entities.RoleBoss) {
		return entities.ErrNotBoss
	}

	err := uc.docRepo.UpdateStatus(ctx, newStatus, user.ID)
	if err != nil {
		return err
	}

	return nil
}
