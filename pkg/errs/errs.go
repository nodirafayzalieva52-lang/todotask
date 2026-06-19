package errs

import "net/http"

type AppError struct {
	HTTPStatus int    `json:"-"`       
	Message    string `json:"error"`   
	LogMessage string `json:"-"`       
}

func (e *AppError) Error() string {
	if e.LogMessage != "" {
		return e.LogMessage
	}
	return e.Message
}

var (
	ErrNotFound = &AppError{
		HTTPStatus: http.StatusNotFound,
		Message:    "Resurs not found",
	}
	ErrBadRequest = &AppError{
		HTTPStatus: http.StatusBadRequest,
		Message:    "Error",
	}
	ErrInternal = &AppError{
		HTTPStatus: http.StatusInternalServerError,
		Message:    "Error",
	}
)

func New(status int, msg string, logMsg string) *AppError {
	return &AppError{
		HTTPStatus: status,
		Message:    msg,
		LogMessage: logMsg,
	}
}