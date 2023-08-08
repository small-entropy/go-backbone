package mongo

import (
	constants "github.com/small-entropy/go-backbone/internal/constants/error"

	"github.com/small-entropy/go-backbone/pkg/datatypes/record"
	errors "github.com/small-entropy/go-backbone/pkg/error"

	facade "github.com/small-entropy/go-backbone/third_party/facade/mongo"

	"time"
)

// InsertOne
// Метод вставки в коллекцию документа
func (s *MongoStore[DATA]) InsertOne(data DATA) (record.Record[facade.ObjectID, DATA], error) {
	var err error
	var result *facade.InsertOneResult
	var inserted record.Record[facade.ObjectID, DATA]

	toInsert := &record.Record[facade.ObjectID, DATA]{
		Identifier: facade.NewObjectID(),
		Data:       data,
		CreatedAt:  &facade.Timestamp{T: uint32(time.Now().Unix())},
	}

	if result, err = s.Storage.InsertOne(*s.Context, toInsert); err == nil {
		identifier_field := s.Filter["Identifier"]
		filter := facade.BsonM{
			identifier_field: result.InsertedID,
		}
		if err = s.Store.Storage.FindOne(*s.Context, filter).Decode(&inserted); err != nil {
			err = &errors.StoreError{
				Status:       constants.ErrStoreDecode,
				StorageName:  s.Storage.Name(),
				DatabaseName: s.Storage.Database().Name(),
				Err:          err,
			}
		}
	} else {
		err = &errors.StoreError{
			Status:       constants.ErrStoreInsert,
			StorageName:  s.Storage.Name(),
			DatabaseName: s.Storage.Database().Name(),
			Err:          err,
		}
	}
	return inserted, err
}
