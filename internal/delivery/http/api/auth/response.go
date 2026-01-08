package auth

import "workflow_engine/internal/domain/entities"

type UserResponse struct {
	ID        int64             `json:"id"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	Phone     string            `json:"phone"`
	Role      entities.UserRole `json:"role"`
	Token     string            `json:"token"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}
