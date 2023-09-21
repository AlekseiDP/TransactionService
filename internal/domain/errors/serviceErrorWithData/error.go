package serviceErrorWithData

// ServiceError Структура для обертки ошибок с доп датой для фронта
type ServiceError struct {
	Data             any
	Err              error
	Message          string
	DeveloperMessage string
	Code             string
}

func (e *ServiceError) Error() string {
	return e.Message
}

func (e *ServiceError) Unwrap() error {
	return e.Err
}

func NewServiceError(err error, message, developerMessage, code string) *ServiceError {
	return &ServiceError{
		Err:              err,
		Message:          message,
		DeveloperMessage: developerMessage,
		Code:             code,
	}
}
