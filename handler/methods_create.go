package handler

import (
	"net/http"

	error_constants "github.com/small-entropy/go-backbone/constants/error"
	"github.com/small-entropy/go-backbone/datatypes/record"
	backbone_error "github.com/small-entropy/go-backbone/error"
	store_provider "github.com/small-entropy/go-backbone/providers/store"
	"github.com/small-entropy/go-backbone/response/jsend"

	"github.com/labstack/echo/v4"
)

// Обработчик создания записи
func (h *Handler[CONN, ID, DATA, DTO]) Create(c echo.Context) error {
	var err error
	var response *jsend.Response
	var dto DTO

	code := http.StatusCreated
	field := h.Settings.Fields.Response["Entity"]

	ctx, cancel := h.GetRequestContext(&c)
	defer cancel()

	if dto, err = h.GetDTO(&c); err == nil {
		var data DATA
		var result record.Record[ID, DATA]
		if data, err = h.Callbacks.Fill(dto); err == nil {
			var provider store_provider.StoreProvider[CONN, ID, DATA]
			if provider, err = h.Callbacks.GetProvider(ctx, h.Settings.StorageName); err == nil {
				if result, err = h.Settings.Controller.InsertOne(data, &provider); err == nil {
					if h.Responses.CreateOne != nil {
						response = h.Responses.CreateOne(&ctx, h, &result)
					} else {
						var data interface{}
						if h.Adapter != nil {
							data = h.Adapter.One(&result)
						} else {
							data = result
						}
						response = jsend.Success(&echo.Map{field: data})
					}
				} else {
					code = http.StatusBadRequest
					err = &backbone_error.HandlerError{
						Status:  error_constants.ERR_HANDLER_CREATE,
						Code:    code,
						Message: error_constants.MSG_HANDLER_CREATE,
						Err:     err,
					}
					response = jsend.Fail(&echo.Map{field: err.Error()})
				}
			} else {
				code = http.StatusInternalServerError
				err = &backbone_error.HandlerError{
					Status:  error_constants.ERR_HANDLER_PROVIDER,
					Code:    code,
					Message: error_constants.MSG_HANDLER_PROVIDER,
					Err:     err,
				}
				response = jsend.Error(error_constants.MSG_HANDLER_PROVIDER, &echo.Map{field: err.Error()}, code)
			}
		} else {
			code = http.StatusBadRequest
			err = &backbone_error.HandlerError{
				Status:  error_constants.ERR_HANDLER_FILL,
				Code:    code,
				Message: error_constants.MSG_HANDLER_FILL,
				Err:     err,
			}
			response = jsend.Fail(&echo.Map{field: err.Error()})
		}
	} else {
		code = http.StatusBadRequest
		err = &backbone_error.HandlerError{
			Status:  error_constants.ERR_HANDLER_DTO,
			Code:    code,
			Message: error_constants.MSG_HANDLER_DTO,
			Err:     err,
		}
		response = jsend.Error(error_constants.MSG_HANDLER_DTO, &echo.Map{field: err.Error()}, code)
	}
	return c.JSON(code, response)
}
