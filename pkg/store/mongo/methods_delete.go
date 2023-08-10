package mongo

import (
	constants "github.com/small-entropy/go-backbone/internal/constants/error"

	errors "github.com/small-entropy/go-backbone/pkg/error"

	facade "github.com/small-entropy/go-backbone/third_party/facade/mongo"

	"github.com/small-entropy/go-backbone/tools/convert"
)

// DeleteOne
// Метод удаления записи из коллекции
func (s *Store[DATA]) DeleteOne(filter map[string]interface{}) (Document[DATA], error) {
	var err error
	var result Document[DATA]

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
func (s *Store[DATA]) DeleteMany(filter map[string]interface{}) (DocumentSet[DATA], error) {
	var err error
	var results DocumentSet[DATA]
	var records []Document[DATA]

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
