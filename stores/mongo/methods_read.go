package mongo

import (
	error_constants "github.com/small-entropy/go-backbone/constants/error"
	"github.com/small-entropy/go-backbone/datatypes/record"
	"github.com/small-entropy/go-backbone/datatypes/recordset"
	backbone_error "github.com/small-entropy/go-backbone/error"
	mongo_facade "github.com/small-entropy/go-backbone/facades/mongo"
	"github.com/small-entropy/go-backbone/stores/abstract"
	"github.com/small-entropy/go-backbone/utils/convert"
)

// FindAll
// Метод поиска получения списка документов из коллекции
func (s *MongoStore[DATA]) FindAll(page abstract.Page, filter map[string]interface{}) (recordset.RecordSet[mongo_facade.ObjectID, DATA], error) {
	var err error
	var cursor *mongo_facade.Cursor
	var results recordset.RecordSet[mongo_facade.ObjectID, DATA]
	var records []record.Record[mongo_facade.ObjectID, DATA]

	results.Meta.Filter = filter
	results.Meta.Limit = page.Limit
	results.Meta.Skip = page.Skip
	// TODO: добавить сортировку
	opts := mongo_facade.GetFindOptions().SetSort(mongo_facade.BsonD{}).SetSkip(page.Skip).SetLimit(page.Limit)

	filter_bson := convert.MapToBsonM(filter)

	if cursor, err = s.Storage.Find(*s.Context, filter_bson, opts); err == nil {
		defer cursor.Close(*s.Context)
		if err = cursor.All(*s.Context, &records); err == nil {
			results.SetItems(records)
		} else {
			err = &backbone_error.StoreError{
				Status:       error_constants.ERR_STORE_DECODE,
				StorageName:  s.Storage.Name(),
				DatabaseName: s.Storage.Database().Name(),
				Err:          err,
			}
		}
	} else {
		err = &backbone_error.StoreError{
			Status:       error_constants.ERR_STORE_READ,
			StorageName:  s.Storage.Name(),
			DatabaseName: s.Storage.Database().Name(),
			Err:          err,
		}
	}
	return results, err
}

// FindOne
// Метод получения одного документа из коллекции
func (s *MongoStore[DATA]) FindOne(filter map[string]interface{}) (record.Record[mongo_facade.ObjectID, DATA], error) {
	var err error
	var result record.Record[mongo_facade.ObjectID, DATA]

	filter_bson := convert.MapToBsonM(filter)

	if err = s.Storage.FindOne(*s.Context, filter_bson).Decode(&result); err != nil {
		var err_status string

		switch err {
		case mongo_facade.ErrNoDocuments:
			err_status = error_constants.ERR_STORE_READ
		case mongo_facade.ErrNilDocument:
			err_status = error_constants.ERR_STORE_READ
		default:
			err_status = error_constants.ERR_STORE_UNKNOWN
		}

		err = &backbone_error.StoreError{
			Status:       err_status,
			StorageName:  s.Storage.Name(),
			DatabaseName: s.Storage.Database().Name(),
			Err:          err,
		}
	}
	return result, err
}
