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

// Create
// Обработчик создания записи
func (h *Handler[CONN, ID, DATA, DTO]) Create(c facade.Context) error {
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
			var prov provider.StoreProvider[CONN, ID, DATA]
			if prov, err = h.Callbacks.GetProvider(ctx, h.Settings.StorageName); err == nil {
				if result, err = h.Settings.Controller.InsertOne(data, &prov); err == nil {
					if h.Responses.CreateOne != nil {
						response = h.Responses.CreateOne(&ctx, h, &result)
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
					code = http.StatusBadRequest
					err = &errors.HandlerError{
						Status:  constants.ErrHandlerCreate,
						Code:    code,
						Message: constants.MsgHandlerCreate,
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
		response = jsend.Error(constants.MsgHandlerDto, &facade.Map{field: err.Error()}, code)
	}
	return c.JSON(code, response)
}
