package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validator wraps the go-playground validator
type Validator struct {
	validate *validator.Validate
}

// New creates a new Validator instance
func New() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

// Validate validates a struct and returns formatted error messages
func (v *Validator) Validate(data interface{}) error {
	err := v.validate.Struct(data)
	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	var errMsgs []string
	for _, fieldErr := range validationErrors {
		errMsgs = append(errMsgs, formatFieldError(fieldErr))
	}

	return &ValidationError{Messages: errMsgs}
}

// ValidationError represents validation errors
type ValidationError struct {
	Messages []string
}

func (e *ValidationError) Error() string {
	return strings.Join(e.Messages, "; ")
}

// formatFieldError formats a single field error into a human-readable message
func formatFieldError(fe validator.FieldError) string {
	field := toSnakeCase(fe.Field())

	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, fe.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}

// toSnakeCase converts CamelCase to snake_case
func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}
