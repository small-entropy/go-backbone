package mongo

import (
	constants "github.com/small-entropy/go-backbone/internal/constants/error"
	"github.com/small-entropy/go-backbone/pkg/datatypes/record"
	errors "github.com/small-entropy/go-backbone/pkg/error"
	prov "github.com/small-entropy/go-backbone/pkg/provider/store"
)

// InsertOne
// Метод создания записи
func (c Controller[DATA, ENTITY]) InsertOne(
	data DATA,
	provider *prov.StoreProvider[CONN, ID, DATA],
) (record.Record[ID, DATA], error) {
	var err error
	var result record.Record[ID, DATA]

	if result, err = provider.Store.InsertOne(data); err != nil {
		err = &errors.ControllerError[DATA]{
			Status:  constants.ErrControllerCreate,
			Message: constants.MsgControllerCreate,
			Data:    &data,
			Err:     err,
		}
	}

	return result, err
}
