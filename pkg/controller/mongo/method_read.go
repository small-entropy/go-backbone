package mongo

import (
	"sync"

	constants "github.com/small-entropy/go-backbone/internal/constants/error"
	"github.com/small-entropy/go-backbone/pkg/datatypes/record"
	"github.com/small-entropy/go-backbone/pkg/datatypes/recordset"
	errors "github.com/small-entropy/go-backbone/pkg/error"
	prov "github.com/small-entropy/go-backbone/pkg/provider/store"
	"github.com/small-entropy/go-backbone/pkg/store/abstract"
)

// FindOne
// Метод получения одной записи по фильтрам
func (c *Controller[CONN, ID, DATA]) FindOne(
	identifier ID,
	deleted bool,
	other_filters *map[string]interface{},
	provider *prov.StoreProvider[CONN, ID, DATA],
) (record.Record[ID, DATA], error) {
	var err error
	var result record.Record[ID, DATA]
	filter := map[string]interface{}{
		c.Fields["Identifier"]: identifier,
		c.Fields["DeletedAt"]:  nil,
	}
	if other_filters != nil {
		for k, v := range *other_filters {
			filter[k] = v
		}
	}

	if result, err = provider.Store.FindOne(filter); err != nil {
		err = &errors.ControllerError[DATA]{
			Status:  constants.ErrControllerRead,
			Message: constants.MsgControllerRead,
			Filter:  &filter,
			Err:     err,
		}
	}
	return result, err
}

// FindOneByFilter
// Метод получения одной записи по фильтрам
func (c *Controller[CONN, ID, DATA]) FindOneByFilter(
	filter map[string]interface{},
	provider *prov.StoreProvider[CONN, ID, DATA],
) (record.Record[ID, DATA], error) {
	var err error
	var result record.Record[ID, DATA]

	if result, err = provider.Store.FindOne(filter); err != nil {
		err = &errors.ControllerError[DATA]{
			Status:  constants.ErrControllerRead,
			Message: constants.MsgControllerRead,
			Filter:  &filter,
			Err:     err,
		}
	}
	return result, err
}

// Find
// Метод получения списка записей по фильтрам и пагинации
func (c *Controller[CONN, ID, DATA]) Find(
	filter map[string]interface{},
	page *abstract.Page,
	provider *prov.StoreProvider[CONN, ID, DATA],
) (recordset.RecordSet[ID, DATA], error) {
	var err error
	var result recordset.RecordSet[ID, DATA]
	var total int64

	var wg sync.WaitGroup
	var err_total error
	var err_read error
	wg.Add(2)
	go func() {
		if total, err_total = provider.Store.GetCount(filter); err_total != nil {
			err = &errors.ControllerError[DATA]{
				Status:  constants.ErrControllerTotal,
				Message: constants.MsgControllerTotal,
				Filter:  &filter,
				Err:     err,
			}
		}
		wg.Done()
	}()
	go func() {
		if result, err_read = provider.Store.FindAll(*page, filter); err_read != nil {
			err = &errors.ControllerError[DATA]{
				Status:  constants.ErrControllerRead,
				Message: constants.MsgControllerRead,
				Filter:  &filter,
				Err:     err,
			}
		}
		wg.Done()
	}()
	wg.Wait()
	if err == nil {
		result.Meta.Total = total
	}
	return result, err
}
