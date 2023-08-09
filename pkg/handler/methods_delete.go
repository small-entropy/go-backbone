package handler

import (
	"net/http"

	constants "github.com/small-entropy/go-backbone/internal/constants/error"
	"github.com/small-entropy/go-backbone/pkg/datatypes/record"
	errors "github.com/small-entropy/go-backbone/pkg/error"
	"github.com/small-entropy/go-backbone/pkg/response/jsend"

	facade "github.com/small-entropy/go-backbone/third_party/facade/echo"
)

// DeleteOne
// Обработчик мягкого удаления
func (h *Handler[CONN, ID, DATA, DTO]) DeleteOne(c facade.Context) error {
	var err error
	var result record.Record[ID, DATA]
	var response *jsend.Response

	code := http.StatusOK
	field := h.GetResponseField("Entity")

	var paramField string
	if paramField, err = h.GetParamField("Id"); err == nil {
		var identifier ID
		idStr := c.Param(paramField)
		if identifier, err = h.Callbacks.GetIdentifierFromString(idStr); err == nil {
			ctx, cancel := h.GetRequestContext(&c)
			defer cancel()

			if provider, err := h.Callbacks.GetProvider(ctx, h.Settings.StorageName); err == nil {
				filter := map[string]interface{}{
					h.Settings.Fields.Filter["Identifier"]: identifier,
				}
				if result, err = h.Settings.Controller.DeleteOne(filter, &provider); err == nil {
					if h.Responses.DeleteOne != nil {
						response = h.Responses.DeleteOne(&ctx, h, &result)
					} else {
						var data interface{}
						if h.Adapter != nil {
							data = h.Adapter.One(&result)
						} else {
							data = result
						}
						response = jsend.Success(&facade.Map{field: data})
					}
				} else {
					code = http.StatusInternalServerError
					err = &errors.HandlerError{
						Status:  constants.ErrHandlerDelete,
						Code:    code,
						Message: constants.MsgHandlerDelete,
						Err:     err,
					}
					response = jsend.Error(constants.MsgHandlerDelete, &facade.Map{field: err.Error()}, code)
				}
			} else {
				code = http.StatusInternalServerError
				err = &errors.HandlerError{
					Status:  constants.ErrHandlerProvider,
					Code:    code,
					Message: constants.MsgHandlerProvider,
					Err:     err,
				}
				response = jsend.Error(constants.MsgHandlerProvider, &facade.Map{field: err.Error()}, code)
			}
		} else {
			code = http.StatusBadRequest
			err = &errors.HandlerError{
				Status:  constants.ErrHandlerConvertID,
				Code:    code,
				Message: constants.ErrHandlerConvertID,
				Err:     err,
			}
			response = jsend.Fail(&facade.Map{field: err.Error()})
		}
	} else {
		code = http.StatusBadRequest
		err = &errors.HandlerError{
			Status:  constants.ErrHandlerParams,
			Code:    code,
			Message: constants.MsgHandlerParams,
			Err:     err,
		}
		response = jsend.Fail(&facade.Map{field: err.Error()})
	}
	return c.JSON(code, response)
}

// EraseOne
// Обработчик перманентного удаления
func (h *Handler[CONN, ID, DATA, DTO]) EraseOne(c facade.Context) error {
	var err error
	var response *jsend.Response
	var result record.Record[ID, DATA]

	code := http.StatusOK
	field := h.GetResponseField("Entity")

	var paramField string
	if paramField, err = h.GetParamField("Id"); err == nil {
		var identifier ID
		idStr := c.Param(paramField)
		if identifier, err = h.Callbacks.GetIdentifierFromString(idStr); err == nil {
			ctx, cancel := h.GetRequestContext(&c)
			defer cancel()
			filter := map[string]interface{}{
				h.Settings.Fields.Filter["Identifier"]: identifier,
			}
			if provider, err := h.Callbacks.GetProvider(ctx, h.Settings.StorageName); err == nil {
				if result, err = h.Settings.Controller.EraseOne(filter, &provider); err == nil {
					if h.Responses.EraseOne != nil {
						response = h.Responses.EraseOne(&ctx, h, &result)
					} else {
						var data interface{}
						if h.Adapter != nil {
							data = h.Adapter.One(&result)
						} else {
							data = result
						}
						response = jsend.Success(&facade.Map{field: data})
					}
				} else {
					code = http.StatusInternalServerError
					err = &errors.HandlerError{
						Status:  constants.ErrHandlerErase,
						Code:    code,
						Message: constants.MsgHandlerErase,
						Err:     err,
					}
					response = jsend.Error(constants.MsgHandlerErase, &facade.Map{field: err.Error()}, code)
				}
			} else {
				code = http.StatusInternalServerError
				err = &errors.HandlerError{
					Status:  constants.ErrHandlerProvider,
					Code:    code,
					Message: constants.MsgHandlerProvider,
					Err:     err,
				}
				response = jsend.Error(constants.MsgHandlerProvider, &facade.Map{field: err.Error()}, code)
			}
		} else {
			code = http.StatusBadRequest
			err = &errors.HandlerError{
				Status:  constants.ErrHandlerConvertID,
				Code:    code,
				Message: constants.MsgHandlerConvertID,
				Err:     err,
			}
			response = jsend.Fail(&facade.Map{field: err.Error()})
		}
	} else {
		code = http.StatusBadRequest
		err = &errors.HandlerError{
			Status:  constants.ErrHandlerParams,
			Code:    code,
			Message: constants.MsgHandlerParams,
			Err:     err,
		}
		response = jsend.Fail(&facade.Map{field: err.Error()})
	}
	return c.JSON(code, response)
}
