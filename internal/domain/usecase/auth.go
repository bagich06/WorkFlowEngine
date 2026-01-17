package usecase

import (
	"context"

	"workflow_engine/internal/domain/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	GetByID(ctx context.Context, id int64) (*entities.User, error)
	GetByRole(ctx context.Context, role string) ([]*entities.User, error)
	GetByPhone(ctx context.Context, phone string) (*entities.User, error)
	GetAll(ctx context.Context) ([]*entities.User, error)
	Delete(ctx context.Context, id int64) error
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}

type AuthUseCase struct {
	userRepo       UserRepository
	passwordHasher PasswordHasher
}

func NewAuthUseCase(userRepo UserRepository, passwordHasher PasswordHasher) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
	}
}

func (uc *AuthUseCase) Register(ctx context.Context, user *entities.User) (*entities.User, error) {
	existingUser, err := uc.userRepo.GetByPhone(ctx, user.Phone)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, entities.ErrUserAlreadyExists
	}

	hashedPassword, err := uc.passwordHasher.Hash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	user.Role = entities.RoleWorker
	createdUser, err := uc.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (uc *AuthUseCase) Login(ctx context.Context, phone, password string) (*entities.User, error) {
	user, err := uc.userRepo.GetByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, entities.ErrUserNotFound
	}

	err = uc.passwordHasher.Compare(user.Password, password)
	if err != nil {
		return nil, entities.ErrInvalidPassword
	}

	return user, nil
}
