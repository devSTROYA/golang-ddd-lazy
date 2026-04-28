package types

type AppError struct {
	Code    string
	Details []FieldError
}

func (e *AppError) Error() string {
	return e.Code
}

func NewValidationError(details []FieldError) *AppError {
	return &AppError{
		Code:    "VALIDATION_FAILED",
		Details: details,
	}
}

func NewDomainError(code string) *AppError {
	return &AppError{
		Code: code,
	}
}
