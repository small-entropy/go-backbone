package handler

import (
	"net/http"

	constants "github.com/small-entropy/go-backbone/internal/constants/error"
	"github.com/small-entropy/go-backbone/pkg/datatypes/record"
	"github.com/small-entropy/go-backbone/pkg/datatypes/recordset"
	errors "github.com/small-entropy/go-backbone/pkg/error"
	provider "github.com/small-entropy/go-backbone/pkg/provider/store"
	"github.com/small-entropy/go-backbone/pkg/response/jsend"
	"github.com/small-entropy/go-backbone/pkg/store/abstract"

	facade "github.com/small-entropy/go-backbone/third_party/facade/echo"
)

// Get
// Обработчик получения одной записи
func (h *Handler[CONN, ID, DATA, DTO]) Get(c facade.Context) error {
	var err error
	var result record.Record[ID, DATA]
	var response *jsend.Response

	code := http.StatusOK
	field := h.Settings.Fields.Response["Entity"]

	ctx, cancel := h.GetRequestContext(&c)
	defer cancel()

	var paramKey string
	if paramKey, err = h.GetParamField("Id"); err == nil {
		paramValue := c.Param(paramKey)
		var identifier ID
		if identifier, err = h.Callbacks.GetIdentifierFromString(paramValue); err == nil {
			if prov, err := h.Callbacks.GetProvider(ctx, h.Settings.StorageName); err == nil {
				if result, err = h.Settings.Controller.FindOne(identifier, false, nil, &prov); err == nil {

					if h.Adapter != nil {
						response = jsend.Success(&facade.Map{field: h.Adapter.One(&result)})
					} else {
						response = jsend.Success(&facade.Map{field: result})
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

// List
// Обработчик получения списка записей
func (h *Handler[CONN, ID, DATA, DTO]) List(c facade.Context) error {
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

	fieldEntity := h.GetResponseField("Entities")

	var prov provider.StoreProvider[CONN, ID, DATA]
	if prov, err = h.Callbacks.GetProvider(ctx, h.Settings.StorageName); err == nil {
		results, err = h.Settings.Controller.Find(filter, &page, &prov)

		fieldMeta := h.GetResponseField("Meta")

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
				response = jsend.Success(&facade.Map{
					fieldEntity: data,
					fieldMeta:   results.Meta,
				})
			}
		} else {
			code = http.StatusBadRequest
			err = &errors.HandlerError{
				Status:  constants.ErrHandlerCreate,
				Code:    code,
				Message: constants.MsgHandlerCreate,
				Err:     err,
			}
			response = jsend.Fail(&facade.Map{fieldEntity: err.Error()})

		}
	} else {
		code = http.StatusInternalServerError
		err = &errors.HandlerError{
			Status:  constants.ErrHandlerProvider,
			Code:    code,
			Message: constants.MsgHandlerProvider,
			Err:     err,
		}
		response = jsend.Error(constants.MsgHandlerProvider, &facade.Map{fieldEntity: err.Error()}, code)
	}
	return c.JSON(code, response)
}
