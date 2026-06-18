package errs

import "net/http"

// AppError — наша кастомная структура ошибки
type AppError struct {
	HTTPStatus int    `json:"-"`       // Статус ответа (не выводится в JSON)
	Message    string `json:"error"`   // Красивое сообщение для клиента
	LogMessage string `json:"-"`       // Подробное сообщение для логгера (техническое)
}

// Реализуем стандартный интерфейс error для Go
func (e *AppError) Error() string {
	if e.LogMessage != "" {
		return e.LogMessage
	}
	return e.Message
}

// Готовые заготовки частых ошибок проекта
var (
	ErrNotFound = &AppError{
		HTTPStatus: http.StatusNotFound,
		Message:    "Ресурс не найден",
	}
	ErrBadRequest = &AppError{
		HTTPStatus: http.StatusBadRequest,
		Message:    "Некорректный запрос или неверный формат данных",
	}
	ErrInternal = &AppError{
		HTTPStatus: http.StatusInternalServerError,
		Message:    "Внутренняя ошибка сервера",
	}
)

// Новый конструктор, если нужно создать уникальную ошибку на лету
func New(status int, msg string, logMsg string) *AppError {
	return &AppError{
		HTTPStatus: status,
		Message:    msg,
		LogMessage: logMsg,
	}
}