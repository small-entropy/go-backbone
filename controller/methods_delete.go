package controller

import (
	"log"

	error_constants "github.com/small-entropy/go-backbone/constants/error"
	"github.com/small-entropy/go-backbone/datatypes/record"
	backbone_error "github.com/small-entropy/go-backbone/error"
	"github.com/small-entropy/go-backbone/providers/store"
)

// Метод мягкого удаления
func (c *Controller[CONN, ID, DATA]) DeleteOne(
	filter map[string]interface{},
	provider *store.StoreProvider[CONN, ID, DATA],
) (record.Record[ID, DATA], error) {
	var err error
	var result record.Record[ID, DATA]

	deletedAt := c.GetTimeNow()
	update := map[string]interface{}{
		c.Fields["DeletedAt"]: deletedAt,
	}
	if result, err = provider.Store.UpdateOne(filter, update); err != nil {
		err = &backbone_error.ControllerError[DATA]{
			Status:  error_constants.ERR_CONTROLLER_DELETE,
			Message: error_constants.MSG_CONTROLLER_DELETE,
			Filter:  &filter,
			Update:  &update,
			Err:     err,
		}
	}
	log.Println(result)
	return result, err
}

// Метод перманентного удаления
func (c *Controller[CONN, ID, DATA]) EraseOne(
	filter map[string]interface{},
	provider *store.StoreProvider[CONN, ID, DATA],
) (record.Record[ID, DATA], error) {
	var err error
	var result record.Record[ID, DATA]
	if result, err = provider.Store.DeleteOne(filter); err != nil {
		err = &backbone_error.ControllerError[DATA]{
			Status:  error_constants.ERR_CONTROLLER_ERASE,
			Message: error_constants.MSG_CONTROLLER_ERASE,
			Filter:  &filter,
			Err:     err,
		}
	}
	return result, err
}
