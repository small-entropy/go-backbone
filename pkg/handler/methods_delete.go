package handler

import (
	"net/http"

	error_constants "github.com/small-entropy/go-backbone/constants/error"
	"github.com/small-entropy/go-backbone/datatypes/record"
	backbone_error "github.com/small-entropy/go-backbone/error"
	"github.com/small-entropy/go-backbone/response/jsend"

	echo_facade "github.com/small-entropy/go-backbone/facade/echo"
)

// DeleteOne
// Обработчик мягкого удаления
func (h *Handler[CONN, ID, DATA, DTO]) DeleteOne(c echo_facade.Context) error {
	var err error
	var result record.Record[ID, DATA]
	var response *jsend.Response

	code := http.StatusOK
	field := h.GetResponseField("Entity")

	var param_field string
	if param_field, err = h.GetParamField("Id"); err == nil {
		var identifier ID
		id_str := c.Param(param_field)
		if identifier, err = h.Callbacks.GetIdentifierFromString(id_str); err == nil {
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
						response = jsend.Success(&echo_facade.Map{field: data})
					}
				} else {
					code = http.StatusInternalServerError
					err = &backbone_error.HandlerError{
						Status:  error_constants.ERR_HANDLER_DELETE,
						Code:    code,
						Message: error_constants.MSG_HANDLER_DELETE,
						Err:     err,
					}
					response = jsend.Error(error_constants.MSG_HANDLER_DELETE, &echo_facade.Map{field: err.Error()}, code)
				}
			} else {
				code = http.StatusInternalServerError
				err = &backbone_error.HandlerError{
					Status:  error_constants.ERR_HANDLER_PROVIDER,
					Code:    code,
					Message: error_constants.MSG_HANDLER_PROVIDER,
					Err:     err,
				}
				response = jsend.Error(error_constants.MSG_HANDLER_PROVIDER, &echo_facade.Map{field: err.Error()}, code)
			}
		} else {
			code = http.StatusBadRequest
			err = &backbone_error.HandlerError{
				Status:  error_constants.ERR_HANDLER_CONVERT_ID,
				Code:    code,
				Message: error_constants.MSG_HANDLER_CONVERT_ID,
				Err:     err,
			}
			response = jsend.Fail(&echo_facade.Map{field: err.Error()})
		}
	} else {
		code = http.StatusBadRequest
		err = &backbone_error.HandlerError{
			Status:  error_constants.ERR_HANDLER_PARAMS,
			Code:    code,
			Message: error_constants.MSG_HANDLER_PARAMS,
			Err:     err,
		}
		response = jsend.Fail(&echo_facade.Map{field: err.Error()})
	}
	return c.JSON(code, response)
}

// EraseOne
// Обработчик перманентного удаления
func (h *Handler[CONN, ID, DATA, DTO]) EraseOne(c echo_facade.Context) error {
	var err error
	var response *jsend.Response
	var result record.Record[ID, DATA]

	code := http.StatusOK
	field := h.GetResponseField("Entity")

	var param_field string
	if param_field, err = h.GetParamField("Id"); err == nil {
		var identifier ID
		id_str := c.Param(param_field)
		if identifier, err = h.Callbacks.GetIdentifierFromString(id_str); err == nil {
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
						response = jsend.Success(&echo_facade.Map{field: data})
					}
				} else {
					code = http.StatusInternalServerError
					err = &backbone_error.HandlerError{
						Status:  error_constants.ERR_HANDLER_ERASE,
						Code:    code,
						Message: error_constants.MSG_HANDLER_ERASE,
						Err:     err,
					}
					response = jsend.Error(error_constants.MSG_HANDLER_ERASE, &echo_facade.Map{field: err.Error()}, code)
				}
			} else {
				code = http.StatusInternalServerError
				err = &backbone_error.HandlerError{
					Status:  error_constants.ERR_HANDLER_PROVIDER,
					Code:    code,
					Message: error_constants.MSG_HANDLER_PROVIDER,
					Err:     err,
				}
				response = jsend.Error(error_constants.MSG_HANDLER_PROVIDER, &echo_facade.Map{field: err.Error()}, code)
			}
		} else {
			code = http.StatusBadRequest
			err = &backbone_error.HandlerError{
				Status:  error_constants.ERR_HANDLER_CONVERT_ID,
				Code:    code,
				Message: error_constants.MSG_HANDLER_CONVERT_ID,
				Err:     err,
			}
			response = jsend.Fail(&echo_facade.Map{field: err.Error()})
		}
	} else {
		code = http.StatusBadRequest
		err = &backbone_error.HandlerError{
			Status:  error_constants.ERR_HANDLER_PARAMS,
			Code:    code,
			Message: error_constants.MSG_HANDLER_PARAMS,
			Err:     err,
		}
		response = jsend.Fail(&echo_facade.Map{field: err.Error()})
	}
	return c.JSON(code, response)
}
