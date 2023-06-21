package jsend

import (
	constants "github.com/small-entropy/go-backbone/constants/jsend"

	"github.com/labstack/echo/v4"
)

// Функция возврата ответа успешного выполнения
func Success(data *echo.Map) *Response {
	return &Response{
		Status: constants.SUCCESS,
		Data:   data,
	}
}

// Функция возврата ответа не успешного выполнения
func Fail(data *echo.Map) *Response {
	return &Response{
		Status: constants.FAILS,
		Data:   data,
	}
}

// Функция возврата ошибки
func Error(message string, data *echo.Map, code int) *Response {
	return &Response{
		Status:  constants.ERROR,
		Message: message,
		Data:    data,
		Code:    code,
	}
}

// Функция возврата простой ошибки
func SimpleError(message string) *Response {
	return &Response{
		Status:  constants.ERROR,
		Message: message,
	}
}
