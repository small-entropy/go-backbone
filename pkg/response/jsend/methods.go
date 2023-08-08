package jsend

import (
	constants "github.com/small-entropy/go-backbone/constants/jsend"
	echo_facade "github.com/small-entropy/go-backbone/facade/echo"
)

// Функция возврата ответа успешного выполнения
func Success(data *echo_facade.Map) *Response {
	return &Response{
		Status: constants.SUCCESS,
		Data:   data,
	}
}

// Функция возврата ответа не успешного выполнения
func Fail(data *echo_facade.Map) *Response {
	return &Response{
		Status: constants.FAILS,
		Data:   data,
	}
}

// Функция возврата ошибки
func Error(message string, data *echo_facade.Map, code int) *Response {
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
