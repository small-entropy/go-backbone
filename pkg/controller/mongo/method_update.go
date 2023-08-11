package mongo

import (
	"encoding/json"
	"time"

	constatns "github.com/small-entropy/go-backbone/internal/constants/error"
	"github.com/small-entropy/go-backbone/pkg/datatypes/record"
	errors "github.com/small-entropy/go-backbone/pkg/error"
	prov "github.com/small-entropy/go-backbone/pkg/provider/store"

	facade "github.com/small-entropy/go-backbone/third_party/facade/mongo"
)

// UpdateOne
// Метод обновления одной записи
func (c *Controller[CONN, ID, DATA]) UpdateOne(
	filter map[string]interface{},
	update DATA,
	provider *prov.StoreProvider[CONN, ID, DATA],
) (record.Record[ID, DATA], error) {
	var err error
	var result record.Record[ID, DATA]
	var updateJson []byte
	var updateMap map[string]interface{}
	if updateJson, err = json.Marshal(update); err == nil {
		if err = json.Unmarshal(updateJson, &updateMap); err == nil {
			// TODO: отвязать от MongoDB
			updatedAt := facade.Timestamp{T: uint32(time.Now().Unix())}
			updateMap[c.Fields["UpdatedAt"]] = updatedAt
			if result, err = provider.Store.UpdateOne(filter, updateMap); err != nil {
				err = &errors.ControllerError[DATA]{
					Status:  constatns.ErrControllerUpdate,
					Message: constatns.MsgControllerUpdate,
					Data:    &update,
					Filter:  &filter,
					Err:     err,
				}
			}
		} else {
			err = &errors.ControllerError[DATA]{
				Status:  constatns.ErrControllerUnMarshalData,
				Message: constatns.MsgControllerUnMarshalData,
				Data:    &update,
				Filter:  &filter,
				Err:     err,
			}
		}
	} else {
		err = &errors.ControllerError[DATA]{
			Status:  constatns.ErrControllerMarshalData,
			Message: constatns.MsgControllerMarshalData,
			Data:    &update,
			Filter:  &filter,
			Err:     err,
		}
	}
	return result, err
}
