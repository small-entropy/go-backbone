package jsend

import (
	constants "github.com/small-entropy/go-backbone/internal/constants/jsend"
	facade "github.com/small-entropy/go-backbone/third_party/facade/echo"
)

// Success
// Функция возврата ответа успешного выполнения
func Success(data *facade.Map) *Response {
	return &Response{
		Status: constants.Success,
		Data:   data,
	}
}

// Fail
// Функция возврата ответа не успешного выполнения
func Fail(data *facade.Map) *Response {
	return &Response{
		Status: constants.Fails,
		Data:   data,
	}
}

// Error
// Функция возврата ошибки
func Error(message string, data *facade.Map, code int) *Response {
	return &Response{
		Status:  constants.Error,
		Message: message,
		Data:    data,
		Code:    code,
	}
}

// SimpleError
// Функция возврата простой ошибки
func SimpleError(message string) *Response {
	return &Response{
		Status:  constants.Error,
		Message: message,
	}
}
