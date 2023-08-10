package mongo

import (
	constants "github.com/small-entropy/go-backbone/internal/constants/error"

	facade "github.com/small-entropy/go-backbone/third_party/facade/mongo"

	"github.com/small-entropy/go-backbone/pkg/datatypes/record"
	"github.com/small-entropy/go-backbone/pkg/datatypes/recordset"
	errors "github.com/small-entropy/go-backbone/pkg/error"
	"github.com/small-entropy/go-backbone/pkg/store/abstract"

	"github.com/small-entropy/go-backbone/tools/convert"
)

// FindAll
// Метод поиска получения списка документов из коллекции
func (s *MongoStore[DATA]) FindAll(page abstract.Page, filter map[string]interface{}) (recordset.RecordSet[facade.ObjectID, DATA], error) {
	var err error
	var cursor *facade.Cursor
	var results recordset.RecordSet[facade.ObjectID, DATA]
	var records []record.Record[facade.ObjectID, DATA]

	results.Meta.Filter = filter
	results.Meta.Limit = page.Limit
	results.Meta.Skip = page.Skip
	// TODO: добавить сортировку
	opts := facade.GetFindOptions().SetSort(facade.BsonD{}).SetSkip(page.Skip).SetLimit(page.Limit)

	currentFilter := convert.MapToBsonM(filter)

	if cursor, err = s.Storage.Find(*s.Context, currentFilter, opts); err == nil {
		defer cursor.Close(*s.Context)
		if err = cursor.All(*s.Context, &records); err == nil {
			results.SetItems(records)
		} else {
			err = &errors.StoreError{
				Status:       constants.ErrStoreDecode,
				StorageName:  s.Storage.Name(),
				DatabaseName: s.Storage.Database().Name(),
				Err:          err,
			}
		}
	} else {
		err = &errors.StoreError{
			Status:       constants.ErrStoreRead,
			StorageName:  s.Storage.Name(),
			DatabaseName: s.Storage.Database().Name(),
			Err:          err,
		}
	}
	return results, err
}

// FindOne
// Метод получения одного документа из коллекции
func (s *MongoStore[DATA]) FindOne(filter map[string]interface{}) (record.Record[facade.ObjectID, DATA], error) {
	var err error
	var result record.Record[facade.ObjectID, DATA]

	currentFilter := convert.MapToBsonM(filter)

	if err = s.Storage.FindOne(*s.Context, currentFilter).Decode(&result); err != nil {
		var status string

		switch err {
		case facade.ErrNoDocuments:
			status = constants.ErrStoreRead
		case facade.ErrNilDocument:
			status = constants.ErrStoreRead
		default:
			status = constants.ErrStoreUnknown
		}

		err = &errors.StoreError{
			Status:       status,
			StorageName:  s.Storage.Name(),
			DatabaseName: s.Storage.Database().Name(),
			Err:          err,
		}
	}
	return result, err
}