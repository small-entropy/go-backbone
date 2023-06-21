package controller

import (
	error_constants "github.com/small-entropy/go-backbone/constants/error"
	"github.com/small-entropy/go-backbone/datatypes/record"
	backbone_error "github.com/small-entropy/go-backbone/error"
	store_provider "github.com/small-entropy/go-backbone/providers/store"
)

// Метод создания записи
func (c Controller[CONN, ID, DATA]) InsertOne(
	data DATA,
	provider *store_provider.StoreProvider[CONN, ID, DATA],
) (record.Record[ID, DATA], error) {
	var err error
	var result record.Record[ID, DATA]
	if result, err = provider.Store.InsertOne(data); err != nil {
		err = &backbone_error.ControllerError[DATA]{
			Status:  error_constants.ERR_CONTROLLER_CREATE,
			Message: error_constants.MSG_CONTROLLER_CREATE,
			Data:    &data,
			Err:     err,
		}
	}
	return result, err
}
