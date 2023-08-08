package mongo

import (
	error_constants "github.com/small-entropy/go-backbone/constants/error"
	"github.com/small-entropy/go-backbone/datatypes/record"
	backbone_error "github.com/small-entropy/go-backbone/error"
	mongo_facade "github.com/small-entropy/go-backbone/facades/mongo"
	"github.com/small-entropy/go-backbone/utils/convert"
)

// DeleteOne
// Метод удаления записи из коллекции
func (s *MongoStore[DATA]) DeleteOne(filter map[string]interface{}) (record.Record[mongo_facade.ObjectID, DATA], error) {
	var err error
	var result record.Record[mongo_facade.ObjectID, DATA]

	filter_bson := convert.MapToBsonM(filter)

	if result, err = s.FindOne(filter_bson); err == nil {
		_, err = s.Storage.DeleteOne(*s.Context, filter_bson)
	} else {
		err = &backbone_error.StoreError{
			Status:       error_constants.ERR_STORE_DELETE,
			StorageName:  s.Storage.Name(),
			DatabaseName: s.Storage.Database().Name(),
			Err:          err,
		}
	}

	return result, err
}
