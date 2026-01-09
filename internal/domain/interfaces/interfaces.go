package interfaces

import (
	"context"
	"workflow_engine/internal/domain/entities"
)

// Repository Interfaces
type UserRepository interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	GetByID(ctx context.Context, id int64) (*entities.User, error)
	GetByRole(ctx context.Context, role string) ([]*entities.User, error)
	GetByPhone(ctx context.Context, phone string) (*entities.User, error)
	GetAll(ctx context.Context) ([]*entities.User, error)
	Delete(ctx context.Context, id int64) error
}

type DocumentRepository interface {
	Create(ctx context.Context, document *entities.Document) (*entities.Document, error)
	GetByID(ctx context.Context, id int64) (*entities.Document, error)
	UpdateStatus(ctx context.Context, newStatus entities.DocumentStatus, id int64) error
	GetStatusByID(ctx context.Context, id int64) (entities.DocumentStatus, error)
}

// UseCase Interfaces
type AuthUseCaseInterface interface {
	Register(ctx context.Context, user *entities.User) (*entities.User, error)
	Login(ctx context.Context, phone, password string) (*entities.User, error)
}

type DocumentUseCase interface {
	Create(ctx context.Context, document *entities.Document) (*entities.Document, error)
	GetByID(ctx context.Context, id int64) (*entities.Document, error)
}

type WorkflowUseCase interface {
	Approve(ctx context.Context, docID int64, userID int64) error
	Reject(ctx context.Context, docID int64, userID int64) error
	GetDocumentStatus(ctx context.Context, docID int64) (entities.DocumentStatus, error)
}

// Service Interfaces
type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}
