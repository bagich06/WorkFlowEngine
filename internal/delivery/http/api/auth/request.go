package auth

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Password  string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}
