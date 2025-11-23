package error

import "fmt"

type ValidationErrorCode string

type ValidationError struct {
	Code   ValidationErrorCode `json:"code"`
}

func NewValidationError(code ValidationErrorCode) *ValidationError {
	return &ValidationError{
		Code: code,
	}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %+v", e.Code)
}
