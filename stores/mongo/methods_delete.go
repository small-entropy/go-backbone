package mongo

import (
	error_constants "github.com/small-entropy/go-backbone/constants/error"
	"github.com/small-entropy/go-backbone/datatypes/record"
	backbone_error "github.com/small-entropy/go-backbone/error"
	"github.com/small-entropy/go-backbone/utils/convert"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Метод удаления записи из коллекции
func (s *MongoStore[DATA]) DeleteOne(filter map[string]interface{}) (record.Record[primitive.ObjectID, DATA], error) {
	var err error
	var result record.Record[primitive.ObjectID, DATA]
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
