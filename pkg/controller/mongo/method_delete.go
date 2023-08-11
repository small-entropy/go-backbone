package mongo

import (
	"log"

	constatns "github.com/small-entropy/go-backbone/internal/constants/error"
	"github.com/small-entropy/go-backbone/pkg/datatypes/record"
	errors "github.com/small-entropy/go-backbone/pkg/error"
	"github.com/small-entropy/go-backbone/pkg/provider/store"
)

// DeleteOne
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
		err = &errors.ControllerError[DATA]{
			Status:  constatns.ErrControllerDelete,
			Message: constatns.MsgControllerDelete,
			Filter:  &filter,
			Update:  &update,
			Err:     err,
		}
	}
	log.Println(result)
	return result, err
}

// EraseOne
// Метод перманентного удаления
func (c *Controller[CONN, ID, DATA]) EraseOne(
	filter map[string]interface{},
	provider *store.StoreProvider[CONN, ID, DATA],
) (record.Record[ID, DATA], error) {
	var err error
	var result record.Record[ID, DATA]
	if result, err = provider.Store.DeleteOne(filter); err != nil {
		err = &errors.ControllerError[DATA]{
			Status:  constatns.ErrControllerErase,
			Message: constatns.MsgControllerErase,
			Filter:  &filter,
			Err:     err,
		}
	}
	return result, err
}
