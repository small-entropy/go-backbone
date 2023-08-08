package mongo

import (
	error_constants "github.com/small-entropy/go-backbone/constants/error"
	"github.com/small-entropy/go-backbone/datatypes/record"
	"github.com/small-entropy/go-backbone/datatypes/recordset"
	backbone_error "github.com/small-entropy/go-backbone/error"
	"github.com/small-entropy/go-backbone/stores/abstract"
	"github.com/small-entropy/go-backbone/utils/convert"
	"go.mongodb.org/mongo-driver/mongo"
)

// Метод поиска получения списка документов из коллекции
func (s *MongoStore[DATA]) FindAll(page abstract.Page, filter map[string]interface{}) (recordset.RecordSet[ObjectID, DATA], error) {
	var err error
	var cursor *Cursor
	var results recordset.RecordSet[ObjectID, DATA]
	var records []record.Record[ObjectID, DATA]

	results.Meta.Filter = filter
	results.Meta.Limit = page.Limit
	results.Meta.Skip = page.Skip
	// TODO: добавить сортировку
	opts := GetFindOptions().SetSort(BsonD{}).SetSkip(page.Skip).SetLimit(page.Limit)
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

// Метод получения одного документа из коллекции
func (s *MongoStore[DATA]) FindOne(filter map[string]interface{}) (record.Record[ObjectID, DATA], error) {
	var err error
	var result record.Record[ObjectID, DATA]
	filter_bson := convert.MapToBsonM(filter)
	if err = s.Storage.FindOne(*s.Context, filter_bson).Decode(&result); err != nil {
		var err_status string
		switch err {
		case mongo.ErrNoDocuments:
			err_status = error_constants.ERR_STORE_READ
		case mongo.ErrNilDocument:
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
