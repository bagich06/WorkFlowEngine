package entities

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user with this credentials already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidPassword   = errors.New("invalid password")
)
