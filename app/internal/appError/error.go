package apperror

import "encoding/json"

type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) UnWrap() error {
	return e.Err
}

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}

	return marshal
}

func SystemError(err error) *AppError {
	return NewAppErr(err, "internal system error", err.Error(), "US-000")
}

func NewAppErr(err error, message, developerMessage string, code string) *AppError {
	return &AppError{
		Err:              err,
		Message:          message,
		DeveloperMessage: developerMessage,
		Code:             code,
	}
}
