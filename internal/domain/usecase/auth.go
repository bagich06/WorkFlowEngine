package usecase

import (
	"context"

	"workflow_engine/internal/domain/entities"
	"workflow_engine/internal/domain/interfaces"
)

type AuthUseCase struct {
	userRepo       interfaces.UserRepository
	passwordHasher interfaces.PasswordHasher
}

func NewAuthUseCase(userRepo interfaces.UserRepository, passwordHasher interfaces.PasswordHasher) *AuthUseCase {
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
