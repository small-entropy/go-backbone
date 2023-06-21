package mongo

import (
	error_constants "github.com/small-entropy/go-backbone/constants/error"
	"github.com/small-entropy/go-backbone/datatypes/record"
	backbone_error "github.com/small-entropy/go-backbone/error"
	"github.com/small-entropy/go-backbone/utils/convert"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Метод обновления одного документа в коллекции
func (s *MongoStore[DATA]) UpdateOne(filter map[string]interface{}, update map[string]interface{}) (record.Record[primitive.ObjectID, DATA], error) {
	var err error
	var result record.Record[primitive.ObjectID, DATA]
	filter_bson := convert.MapToBsonM(filter)
	update_bson := convert.MapToBsonM(update)
	update_query := bson.M{
		"$set": update_bson,
	}
	if _, err = s.Storage.UpdateOne(*s.Context, filter_bson, update_query); err == nil {
		err = s.Storage.FindOne(*s.Context, filter_bson).Decode(&result)
	} else {
		err = &backbone_error.StoreError{
			Status:       error_constants.ERR_STORE_UPDATE,
			StorageName:  s.Storage.Name(),
			DatabaseName: s.Storage.Database().Name(),
			Err:          err,
		}
	}
	return result, err
}
