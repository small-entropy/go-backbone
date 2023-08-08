package handler

import (
	"strconv"

	constants "github.com/small-entropy/go-backbone/constants/handler"

	"context"
	"errors"
	"time"

	echo_facade "github.com/small-entropy/go-backbone/facade/echo"
)

// GetParamField
// Метод получения параметра запроса
func (h *Handler[CONN, ID, DATA, DTO]) GetParamField(key string) (string, error) {
	var err error
	param_key := h.Settings.Fields.Params[key]
	if len(param_key) == 0 {
		err = errors.New("can not find url param")
	}
	return param_key, err
}

// GetCtxTimeout
// Метод получения таймаута
func (h *Handler[CONN, ID, DATA, DTO]) GetCtxTimeout() time.Duration {
	timeout := 10 * time.Second
	return timeout
}

// GetRequestContext
// Метод получения контекста выполнения обработчика
func (h *Handler[CONN, ID, DATA, DTO]) GetRequestContext(c *echo_facade.Context) (context.Context, context.CancelFunc) {
	echo_ctx := *c
	timeout := h.GetCtxTimeout()
	ctx, cancel := context.WithTimeout(echo_ctx.Request().Context(), timeout)
	return ctx, cancel
}

// GetDTO
// Метод получения Data Transfer Object
func (h *Handler[CONN, ID, DATA, DTO]) GetDTO(e *echo_facade.Context) (DTO, error) {
	var err error
	var dto DTO

	ctx := *e

	if err = ctx.Bind(&dto); err == nil {
		err = h.Settings.Controller.GetValidator().Struct(&dto)
	}

	return dto, err
}

// GetSkipFromQuery
// Функция получения количество пропускаемых записей
func (h *Handler[CONN, ID, DATA, DTO]) GetSkipFromQuery(c *echo_facade.Context) int64 {
	var skip int64
	var err error
	ctx := *c
	str_skip := ctx.QueryParam(constants.QUERY_SKIP)
	if len(str_skip) > 0 {
		if skip, err = strconv.ParseInt(str_skip, 10, 64); err != nil {
			skip = constants.LIST_SKIP
		}
	} else {
		skip = constants.LIST_SKIP
	}

	return skip
}

// GetLimitFromQuery
// Функция получения максимального количество возвращаемых записей
func (h *Handler[CONN, ID, DATA, DTO]) GetLimitFromQuery(c *echo_facade.Context) int64 {
	var limit int64
	var err error
	ctx := *c
	str_limit := ctx.QueryParam(constants.QUERY_LIMIT)
	if len(str_limit) > 0 {
		if limit, err = strconv.ParseInt(str_limit, 10, 64); err != nil {
			limit = constants.LIST_LIMIT
		}
	} else {
		limit = constants.LIST_LIMIT
	}

	return limit
}

// GetResponseField
// Метод получения поля для ответа
func (h *Handler[CONN, ID, DATA, DTO]) GetResponseField(key string) string {
	response_field := h.Settings.Fields.Response[key]
	if len(response_field) == 0 {
		response_field = "Unknown"
	}
	return response_field
}
