package handler

import (
	"net/http"

	error_constants "github.com/small-entropy/go-backbone/constants/error"
	"github.com/small-entropy/go-backbone/datatypes/record"
	backbone_error "github.com/small-entropy/go-backbone/error"
	echo_facade "github.com/small-entropy/go-backbone/facade/echo"
	store_provider "github.com/small-entropy/go-backbone/providers/store"
	"github.com/small-entropy/go-backbone/response/jsend"
)

// UpdateOne
// Обработчик обновления одной записи
func (h *Handler[CONN, ID, DATA, DTO]) UpdateOne(c echo_facade.Context) error {
	var err error
	var result record.Record[ID, DATA]
	var response *jsend.Response

	code := http.StatusOK
	field := h.Settings.Fields.Response["Entity"]

	var param string
	if param, err = h.GetParamField("Id"); err == nil {
		id_str := c.Param(param)

		ctx, cancel := h.GetRequestContext(&c)
		defer cancel()

		var provider store_provider.StoreProvider[CONN, ID, DATA]
		if provider, err = h.Callbacks.GetProvider(ctx, h.Settings.StorageName); err == nil {
			var identifier ID
			if identifier, err = h.Callbacks.GetIdentifierFromString(id_str); err == nil {
				var dto DTO
				if dto, err = h.GetDTO(&c); err == nil {
					var update DATA
					if update, err = h.Callbacks.Fill(dto); err == nil {
						filter := map[string]interface{}{
							h.Settings.Fields.Filter["Identifier"]: identifier,
						}
						if result, err = h.Settings.Controller.UpdateOne(filter, update, &provider); err == nil {
							if h.Responses.UpdateOne != nil {
								response = h.Responses.UpdateOne(&ctx, h, &result)
							} else {
								var data interface{}
								if h.Adapter != nil {
									data = h.Adapter.One(&result)
								} else {
									data = result
								}
								response = jsend.Success(&echo_facade.Map{
									field: data,
								})
							}
						} else {
							code = http.StatusInternalServerError
							err = &backbone_error.HandlerError{
								Status:  error_constants.ERR_HANDLER_UPDATE,
								Code:    code,
								Message: error_constants.MSG_HANDLER_UPDATE,
								Err:     err,
							}
							response = jsend.Fail(&echo_facade.Map{field: err.Error()})
						}
					} else {
						code = http.StatusBadRequest
						err = &backbone_error.HandlerError{
							Status:  error_constants.ERR_HANDLER_FILL,
							Code:    code,
							Message: error_constants.MSG_HANDLER_FILL,
							Err:     err,
						}
						response = jsend.Fail(&echo_facade.Map{field: err.Error()})
					}
				} else {
					code = http.StatusBadRequest
					err = &backbone_error.HandlerError{
						Status:  error_constants.ERR_HANDLER_DTO,
						Code:    code,
						Message: error_constants.MSG_HANDLER_DTO,
						Err:     err,
					}
					response = jsend.Fail(&echo_facade.Map{field: err.Error()})
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
			Status:  error_constants.ERR_HANDLER_PARAMS,
			Code:    code,
			Message: error_constants.MSG_HANDLER_PARAMS,
			Err:     err,
		}
		response = jsend.Fail(&echo_facade.Map{field: err.Error()})
	}
	return c.JSON(code, response)
}
