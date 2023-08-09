package handler

import (
	"net/http"

	constants "github.com/small-entropy/go-backbone/internal/constants/error"
	"github.com/small-entropy/go-backbone/pkg/datatypes/record"
	errors "github.com/small-entropy/go-backbone/pkg/error"
	provider "github.com/small-entropy/go-backbone/pkg/provider/store"
	"github.com/small-entropy/go-backbone/pkg/response/jsend"
	facade "github.com/small-entropy/go-backbone/third_party/facade/echo"
)

// UpdateOne
// Обработчик обновления одной записи
func (h *Handler[CONN, ID, DATA, DTO]) UpdateOne(c facade.Context) error {
	var err error
	var result record.Record[ID, DATA]
	var response *jsend.Response

	code := http.StatusOK
	field := h.Settings.Fields.Response["Entity"]

	var param string
	if param, err = h.GetParamField("Id"); err == nil {
		idStr := c.Param(param)

		ctx, cancel := h.GetRequestContext(&c)
		defer cancel()

		var prov provider.StoreProvider[CONN, ID, DATA]
		if prov, err = h.Callbacks.GetProvider(ctx, h.Settings.StorageName); err == nil {
			var identifier ID
			if identifier, err = h.Callbacks.GetIdentifierFromString(idStr); err == nil {
				var dto DTO
				if dto, err = h.GetDTO(&c); err == nil {
					var update DATA
					if update, err = h.Callbacks.Fill(dto); err == nil {
						filter := map[string]interface{}{
							h.Settings.Fields.Filter["Identifier"]: identifier,
						}
						if result, err = h.Settings.Controller.UpdateOne(filter, update, &prov); err == nil {
							if h.Responses.UpdateOne != nil {
								response = h.Responses.UpdateOne(&ctx, h, &result)
							} else {
								var data interface{}
								if h.Adapter != nil {
									data = h.Adapter.One(&result)
								} else {
									data = result
								}
								response = jsend.Success(&facade.Map{
									field: data,
								})
							}
						} else {
							code = http.StatusInternalServerError
							err = &errors.HandlerError{
								Status:  constants.ErrHandlerUpdate,
								Code:    code,
								Message: constants.MsgHandlerUpdate,
								Err:     err,
							}
							response = jsend.Fail(&facade.Map{field: err.Error()})
						}
					} else {
						code = http.StatusBadRequest
						err = &errors.HandlerError{
							Status:  constants.ErrHandlerFill,
							Code:    code,
							Message: constants.MsgHandlerFill,
							Err:     err,
						}
						response = jsend.Fail(&facade.Map{field: err.Error()})
					}
				} else {
					code = http.StatusBadRequest
					err = &errors.HandlerError{
						Status:  constants.ErrHandlerDto,
						Code:    code,
						Message: constants.MsgHandlerDto,
						Err:     err,
					}
					response = jsend.Fail(&facade.Map{field: err.Error()})
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
			Status:  constants.ErrHandlerParams,
			Code:    code,
			Message: constants.MsgHandlerParams,
			Err:     err,
		}
		response = jsend.Fail(&facade.Map{field: err.Error()})
	}
	return c.JSON(code, response)
}
