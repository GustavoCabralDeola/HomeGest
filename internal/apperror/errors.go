package apperror

import "fmt"

// NotFoundError representa um erro de recurso não encontrado (equivalente ao KeyNotFoundException do C#).
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

// NewNotFound cria um novo NotFoundError.
func NewNotFound(format string, args ...interface{}) *NotFoundError {
	return &NotFoundError{Message: fmt.Sprintf(format, args...)}
}

// ValidationError representa um erro de validação de regra de negócio (equivalente ao Exception do C#).
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// NewValidation cria um novo ValidationError.
func NewValidation(format string, args ...interface{}) *ValidationError {
	return &ValidationError{Message: fmt.Sprintf(format, args...)}
}

// IsNotFound verifica se o erro é do tipo NotFoundError.
func IsNotFound(err error) bool {
	_, ok := err.(*NotFoundError)
	return ok
}

// IsValidation verifica se o erro é do tipo ValidationError.
func IsValidation(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}
