package handler

import (
	"strconv"

	constants "github.com/small-entropy/go-backbone/internal/constants/handler"

	"context"
	"errors"
	"time"

	facade "github.com/small-entropy/go-backbone/third_party/facade/echo"
)

// GetParamField
// Метод получения параметра запроса
func (h *Handler[CONN, ID, DATA, DTO]) GetParamField(key string) (string, error) {
	var err error
	paramKey := h.Settings.Fields.Params[key]
	if len(paramKey) == 0 {
		err = errors.New("can not find url param")
	}
	return paramKey, err
}

// GetCtxTimeout
// Метод получения таймаута
func (h *Handler[CONN, ID, DATA, DTO]) GetCtxTimeout() time.Duration {
	timeout := 10 * time.Second
	return timeout
}

// GetRequestContext
// Метод получения контекста выполнения обработчика
func (h *Handler[CONN, ID, DATA, DTO]) GetRequestContext(c *facade.Context) (context.Context, context.CancelFunc) {
	echo_ctx := *c
	timeout := h.GetCtxTimeout()
	ctx, cancel := context.WithTimeout(echo_ctx.Request().Context(), timeout)
	return ctx, cancel
}

// GetDTO
// Метод получения Data Transfer Object
func (h *Handler[CONN, ID, DATA, DTO]) GetDTO(e *facade.Context) (DTO, error) {
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
func (h *Handler[CONN, ID, DATA, DTO]) GetSkipFromQuery(c *facade.Context) int64 {
	var skip int64
	var err error
	ctx := *c
	skipStr := ctx.QueryParam(constants.QuerySkip)
	if len(skipStr) > 0 {
		if skip, err = strconv.ParseInt(skipStr, 10, 64); err != nil {
			skip = constants.DefaultSkipValue
		}
	} else {
		skip = constants.DefaultSkipValue
	}

	return skip
}

// GetLimitFromQuery
// Функция получения максимального количество возвращаемых записей
func (h *Handler[CONN, ID, DATA, DTO]) GetLimitFromQuery(c *facade.Context) int64 {
	var limit int64
	var err error
	ctx := *c
	limitStr := ctx.QueryParam(constants.QueryLimit)
	if len(limitStr) > 0 {
		if limit, err = strconv.ParseInt(limitStr, 10, 64); err != nil {
			limit = constants.DefaultLimitValue
		}
	} else {
		limit = constants.DefaultLimitValue
	}

	return limit
}

// GetResponseField
// Метод получения поля для ответа
func (h *Handler[CONN, ID, DATA, DTO]) GetResponseField(key string) string {
	field := h.Settings.Fields.Response[key]
	if len(field) == 0 {
		field = "Unknown"
	}
	return field
}
