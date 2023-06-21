package controller

import (
	"sync"

	error_constants "github.com/small-entropy/go-backbone/constants/error"
	"github.com/small-entropy/go-backbone/datatypes/record"
	"github.com/small-entropy/go-backbone/datatypes/recordset"
	backbone_error "github.com/small-entropy/go-backbone/error"
	store_provider "github.com/small-entropy/go-backbone/providers/store"
	"github.com/small-entropy/go-backbone/stores/abstract"
)

// Метод получения одной записи по фильтрам
func (c *Controller[CONN, ID, DATA]) FindOne(
	identifier ID,
	deleted bool,
	other_filters *map[string]interface{},
	provider *store_provider.StoreProvider[CONN, ID, DATA],
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
		err = &backbone_error.ControllerError[DATA]{
			Status:  error_constants.ERR_CONTROLLER_READ,
			Message: error_constants.MSG_CONTROLLER_READ,
			Filter:  &filter,
			Err:     err,
		}
	}
	return result, err
}

// Метод получения одной записи по фильтрам
func (c *Controller[CONN, ID, DATA]) FindOneByFilter(
	filter map[string]interface{},
	provider *store_provider.StoreProvider[CONN, ID, DATA],
) (record.Record[ID, DATA], error) {
	var err error
	var result record.Record[ID, DATA]

	if result, err = provider.Store.FindOne(filter); err != nil {
		err = &backbone_error.ControllerError[DATA]{
			Status:  error_constants.ERR_CONTROLLER_READ,
			Message: error_constants.MSG_CONTROLLER_READ,
			Filter:  &filter,
			Err:     err,
		}
	}
	return result, err
}

// Метод получения списка записей по фильтрам и пагинации
func (c *Controller[CONN, ID, DATA]) Find(
	filter map[string]interface{},
	page *abstract.Page,
	provider *store_provider.StoreProvider[CONN, ID, DATA],
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
			err = &backbone_error.ControllerError[DATA]{
				Status:  error_constants.ERR_CONTROLLER_TOTAL,
				Message: error_constants.MSG_CONTROLLER_TOTAL,
				Filter:  &filter,
				Err:     err,
			}
		}
		wg.Done()
	}()
	go func() {
		if result, err_read = provider.Store.FindAll(*page, filter); err_read != nil {
			err = &backbone_error.ControllerError[DATA]{
				Status:  error_constants.ERR_CONTROLLER_READ,
				Message: error_constants.MSG_CONTROLLER_READ,
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
