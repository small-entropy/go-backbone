package controller

import (
	"encoding/json"
	"time"

	error_constants "github.com/small-entropy/go-backbone/constants/error"
	"github.com/small-entropy/go-backbone/datatypes/record"
	backbone_error "github.com/small-entropy/go-backbone/error"
	store_provider "github.com/small-entropy/go-backbone/providers/store"

	mongo_facade "github.com/small-entropy/go-backbone/facade/mongo"
)

// UpdateOne
// Метод обновления одной записи
func (c *Controller[CONN, ID, DATA]) UpdateOne(
	filter map[string]interface{},
	update DATA,
	provider *store_provider.StoreProvider[CONN, ID, DATA],
) (record.Record[ID, DATA], error) {
	var err error
	var result record.Record[ID, DATA]
	var update_json []byte
	var update_map map[string]interface{}
	if update_json, err = json.Marshal(update); err == nil {
		if err = json.Unmarshal(update_json, &update_map); err == nil {
			// TODO: отвязать от MongoDB
			updatedAt := mongo_facade.Timestamp{T: uint32(time.Now().Unix())}
			update_map[c.Fields["UpdatedAt"]] = updatedAt
			if result, err = provider.Store.UpdateOne(filter, update_map); err != nil {
				err = &backbone_error.ControllerError[DATA]{
					Status:  error_constants.ERR_CONTROLLER_UPDATE,
					Message: error_constants.MSG_CONTROLLER_UPDATE,
					Data:    &update,
					Filter:  &filter,
					Err:     err,
				}
			}
		} else {
			err = &backbone_error.ControllerError[DATA]{
				Status:  error_constants.ERR_CONTROLLER_UNMARSHAL_DATA,
				Message: error_constants.MSG_CONTROLLER_UNMARSHAL_DATA,
				Data:    &update,
				Filter:  &filter,
				Err:     err,
			}
		}
	} else {
		err = &backbone_error.ControllerError[DATA]{
			Status:  error_constants.ERR_CONTROLLER_MARSHAL_DATA,
			Message: error_constants.MSG_CONTROLLER_MARSHAL_DATA,
			Data:    &update,
			Filter:  &filter,
			Err:     err,
		}
	}
	return result, err
}
