package mongo

import (
	constants "github.com/small-entropy/go-backbone/internal/constants/error"

	"github.com/small-entropy/go-backbone/pkg/datatypes/record"
	"github.com/small-entropy/go-backbone/pkg/datatypes/recordset"
	errors "github.com/small-entropy/go-backbone/pkg/error"

	facade "github.com/small-entropy/go-backbone/third_party/facade/mongo"

	"github.com/small-entropy/go-backbone/tools/convert"
)

// DeleteOne
// Метод удаления записи из коллекции
func (s *MongoStore[DATA]) DeleteOne(filter map[string]interface{}) (record.Record[facade.ObjectID, DATA], error) {
	var err error
	var result record.Record[facade.ObjectID, DATA]

	currentFilter := convert.MapToBsonM(filter)

	if result, err = s.FindOne(currentFilter); err == nil {
		_, err = s.Storage.DeleteOne(*s.Context, currentFilter)
	} else {
		err = &errors.StoreError{
			Status:       constants.ErrStoreDelete,
			StorageName:  s.Storage.Name(),
			DatabaseName: s.Storage.Database().Name(),
			Err:          err,
		}
	}

	return result, err
}

// DeleteMany
// Метод удаления списка записей из коллекции по фильтру
func (s *MongoStore[DATA]) DeleteMany(filter map[string]interface{}) (recordset.RecordSet[facade.ObjectID, DATA], error) {
	var err error
	var results recordset.RecordSet[facade.ObjectID, DATA]
	var records []record.Record[facade.ObjectID, DATA]

	currentFilter := convert.MapToBsonM(filter)

	var cursor *facade.Cursor
	if cursor, err = s.Storage.Find(*s.Context, currentFilter); err == nil {
		defer cursor.Close(*s.Context)
		if err = cursor.All(*s.Context, &records); err == nil {
			results.SetItems(records)
			_, err = s.DeleteMany(currentFilter)
		} else {
			err = &errors.StoreError{
				Status:       constants.ErrStoreDecode,
				StorageName:  s.Storage.Name(),
				DatabaseName: s.Storage.Database().Name(),
				Err:          err,
			}
		}
	}

	return results, err
}
