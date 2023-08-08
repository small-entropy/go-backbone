package handler

import (
	"net/http"

	error_constants "github.com/small-entropy/go-backbone/constants/error"
	"github.com/small-entropy/go-backbone/datatypes/record"
	"github.com/small-entropy/go-backbone/datatypes/recordset"
	backbone_error "github.com/small-entropy/go-backbone/error"
	store_provider "github.com/small-entropy/go-backbone/providers/store"
	"github.com/small-entropy/go-backbone/response/jsend"
	"github.com/small-entropy/go-backbone/stores/abstract"

	echo_facade "github.com/small-entropy/go-backbone/facades/echo"
)

// Get
// Обработчик получения одной записи
func (h *Handler[CONN, ID, DATA, DTO]) Get(c echo_facade.Context) error {
	var err error
	var result record.Record[ID, DATA]
	var response *jsend.Response

	code := http.StatusOK
	field := h.Settings.Fields.Response["Entity"]

	ctx, cancel := h.GetRequestContext(&c)
	defer cancel()

	var param_key string
	if param_key, err = h.GetParamField("Id"); err == nil {
		id_param_value := c.Param(param_key)
		var identifier ID
		if identifier, err = h.Callbacks.GetIdentifierFromString(id_param_value); err == nil {
			if provider, err := h.Callbacks.GetProvider(ctx, h.Settings.StorageName); err == nil {
				if result, err = h.Settings.Controller.FindOne(identifier, false, nil, &provider); err == nil {

					if h.Adapter != nil {
						response = jsend.Success(&echo_facade.Map{field: h.Adapter.One(&result)})
					} else {
						response = jsend.Success(&echo_facade.Map{field: result})
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

// List
// Обработчик получения списка записей
func (h *Handler[CONN, ID, DATA, DTO]) List(c echo_facade.Context) error {
	var err error
	var response *jsend.Response
	var results recordset.RecordSet[ID, DATA]

	code := http.StatusOK

	page := abstract.Page{
		Limit: h.GetLimitFromQuery(&c),
		Skip:  h.GetSkipFromQuery(&c),
	}

	ctx, cancel := h.GetRequestContext(&c)
	defer cancel()

	filter := map[string]interface{}{
		h.Settings.Fields.Filter["DeletedAt"]: nil,
	}

	field_entities := h.GetResponseField("Entities")

	var provider store_provider.StoreProvider[CONN, ID, DATA]
	if provider, err = h.Callbacks.GetProvider(ctx, h.Settings.StorageName); err == nil {
		results, err = h.Settings.Controller.Find(filter, &page, &provider)

		field_meta := h.GetResponseField("Meta")

		if err == nil {
			if h.Responses.List != nil {
				response = h.Responses.List(&ctx, h, &results)
			} else {
				var data interface{}
				if h.Adapter != nil {
					data = h.Adapter.List(&results)
				} else {
					data = results
				}
				response = jsend.Success(&echo_facade.Map{
					field_entities: data,
					field_meta:     results.Meta,
				})
			}
		} else {
			code = http.StatusBadRequest
			err = &backbone_error.HandlerError{
				Status:  error_constants.ERR_HANDLER_CREATE,
				Code:    code,
				Message: error_constants.MSG_HANDLER_CREATE,
				Err:     err,
			}
			response = jsend.Fail(&echo_facade.Map{field_entities: err.Error()})

		}
	} else {
		code = http.StatusInternalServerError
		err = &backbone_error.HandlerError{
			Status:  error_constants.ERR_HANDLER_PROVIDER,
			Code:    code,
			Message: error_constants.MSG_HANDLER_PROVIDER,
			Err:     err,
		}
		response = jsend.Error(error_constants.MSG_HANDLER_PROVIDER, &echo_facade.Map{field_entities: err.Error()}, code)
	}
	return c.JSON(code, response)
}
