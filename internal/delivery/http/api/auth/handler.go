package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"workflow_engine/internal/domain/entities"
	"workflow_engine/pkg/jwt"
	"workflow_engine/pkg/validator"
)

type AuthUseCaseInterface interface {
	Register(ctx context.Context, user *entities.User) (*entities.User, error)
	Login(ctx context.Context, phone, password string) (*entities.User, error)
}

type AuthHandler struct {
	authUseCase AuthUseCaseInterface
	jwtService  jwt.JWTServiceInterface
	validator   *validator.Validator
}

func NewAuthHandler(authUseCase AuthUseCaseInterface, jwtService jwt.JWTServiceInterface) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
		jwtService:  jwtService,
		validator:   validator.New(),
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.validator.Validate(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user := &entities.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Password:  req.Password,
	}

	createdUser, err := h.authUseCase.Register(r.Context(), user)
	if err != nil {
		if errors.Is(err, entities.ErrUserAlreadyExists) {
			h.respondWithError(w, http.StatusConflict, "User with this credentials already exists")
			return
		}
		h.respondWithError(w, http.StatusInternalServerError, "Failed to register student")
		return
	}

	token, err := h.jwtService.GenerateToken(createdUser.ID, createdUser.Role)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	response := UserResponse{
		ID:        createdUser.ID,
		FirstName: createdUser.FirstName,
		LastName:  createdUser.LastName,
		Phone:     createdUser.Phone,
		Role:      createdUser.Role,
		Token:     token,
	}

	h.respondWithJSON(w, http.StatusCreated, response)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.validator.Validate(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.authUseCase.Login(r.Context(), req.Phone, req.Password)
	if err != nil {
		if errors.Is(err, entities.ErrUserAlreadyExists) {
			h.respondWithError(w, http.StatusUnauthorized, "Invalid phone or password")
			return
		}
		h.respondWithError(w, http.StatusInternalServerError, "Failed to login")
		return
	}

	token, err := h.jwtService.GenerateToken(user.ID, user.Role)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	response := UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		Role:      user.Role,
		Token:     token,
	}

	h.respondWithJSON(w, http.StatusOK, response)
}

func (h *AuthHandler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func (h *AuthHandler) respondWithError(w http.ResponseWriter, code int, message string) {
	h.respondWithJSON(w, code, ErrorResponse{
		Error:   http.StatusText(code),
		Message: message,
	})
}
