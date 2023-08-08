package mongo

import (
	error_constants "github.com/small-entropy/go-backbone/constants/error"
	"github.com/small-entropy/go-backbone/datatypes/record"
	backbone_error "github.com/small-entropy/go-backbone/error"
	mongo_facade "github.com/small-entropy/go-backbone/facades/mongo"

	"time"
)

// InsertOne
// Метод вставки в коллекцию документа
func (s *MongoStore[DATA]) InsertOne(data DATA) (record.Record[mongo_facade.ObjectID, DATA], error) {
	var err error
	var result *mongo_facade.InsertOneResult
	var inserted record.Record[mongo_facade.ObjectID, DATA]

	to_insert := &record.Record[mongo_facade.ObjectID, DATA]{
		Identifier: mongo_facade.NewObjectID(),
		Data:       data,
		CreatedAt:  &mongo_facade.Timestamp{T: uint32(time.Now().Unix())},
	}

	if result, err = s.Storage.InsertOne(*s.Context, to_insert); err == nil {
		identifier_field := s.Filter["Identifier"]
		filter := mongo_facade.BsonM{
			identifier_field: result.InsertedID,
		}
		if err = s.Store.Storage.FindOne(*s.Context, filter).Decode(&inserted); err != nil {
			err = &backbone_error.StoreError{
				Status:       error_constants.ERR_STORE_DECODE,
				StorageName:  s.Storage.Name(),
				DatabaseName: s.Storage.Database().Name(),
				Err:          err,
			}
		}
	} else {
		err = &backbone_error.StoreError{
			Status:       error_constants.ERR_STORE_INSERT,
			StorageName:  s.Storage.Name(),
			DatabaseName: s.Storage.Database().Name(),
			Err:          err,
		}
	}
	return inserted, err
}
