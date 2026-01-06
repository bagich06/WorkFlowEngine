package entities

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user with this credentials already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidPassword   = errors.New("invalid password")

	ErrNotWorker    = errors.New("only workers are allowed to do this")
	ErrNotManager   = errors.New("only managers are allowed to do this")
	ErrNotEconomist = errors.New("only economists is allowed to do this")
	ErrNotBoss      = errors.New("only boss is allowed to do this")
)
