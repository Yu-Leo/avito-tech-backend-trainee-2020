package apperror

var (
	UsernameAlreadyExists = NewAppError(
		"The user with the specified username already exists",
		"")

	IDNotFound = NewAppError(
		"The object with the specified id not found",
		"")

	InternalServerError = NewAppError(
		"Internal Server Error",
		"")

	ChatNameAlreadyExists = NewAppError(
		"The chat with the specified name already exists",
		"")

	UserIsNotInChat = NewAppError(
		"The author of the message is not a member of the chat",
		"")

	TooLongName = NewAppError(
		"The name must be shorter than 80 characters",
		"")
)

type ErrorJSON struct {
	Message string `json:"message"`
}

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
