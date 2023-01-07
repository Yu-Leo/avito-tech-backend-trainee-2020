package apperror

var (
	UsernameAlreadyExists = NewAppError(
		"The user with the specified username already exists",
		"")

	UserIDNotFound = NewAppError(
		"The user with the specified id not found",
		"")
)

type AppError struct {
	Message          string `json:"message"`
	DeveloperMessage string `json:"developerMessage"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(message, developerMessage string) *AppError {
	return &AppError{
		Message:          message,
		DeveloperMessage: developerMessage,
	}
}
