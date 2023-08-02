package mongo

import (
	error_constants "github.com/small-entropy/go-backbone/constants/error"
	"github.com/small-entropy/go-backbone/datatypes/record"
	"github.com/small-entropy/go-backbone/datatypes/recordset"
	backbone_error "github.com/small-entropy/go-backbone/error"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertMany
// Метод вставки в коллекцию нескольких документов
func (s *MongoStore[DATA]) InsertMany(data []DATA) (recordset.RecordSet[primitive.ObjectID, DATA], error) {
	var err error
	var results recordset.RecordSet[primitive.ObjectID, DATA]
	// Объявляем массив документов для вставки
	var to_insert []interface{}
	// Получаем время, которое будет считаться временем создания всех документов
	createdAt := &primitive.Timestamp{T: uint32(time.Now().Unix())}
	// Создаем массив идентификаторов фиксированной длины
	oids := make([]primitive.ObjectID, len(to_insert))
	// Объявляем переменную для идентификатора документа
	var oid primitive.ObjectID
	// Перебираем данные для вставки
	for _, raw := range data {
		/// Получаем новый уникальный идентификатор
		oid = primitive.NewObjectID()
		/// Добавляем идентификатор в массив
		oids = append(oids, oid)
		/// Собираем Record документа
		to_insert = append(to_insert, record.Record[primitive.ObjectID, DATA]{
			Identifier: oid,
			Data:       raw,
			CreatedAt:  createdAt,
		})
	}

	// Пытаемся вставить документы
	if _, err = s.Storage.InsertMany(*s.Context, to_insert); err == nil {
		/// Если документы успешно вставлены, то собираем фильтр.
		/// Получаем наименование поля-идентификатора в фильтре
		identifier_field := s.Filter["Identifier"]
		/// Собираем фильтр
		filter := bson.M{
			identifier_field: bson.M{
				"$in": oids,
			},
		}
		/// Объявляем курсор и список Record'ов
		var cursor *mongo.Cursor
		var records []record.Record[primitive.ObjectID, DATA]
		/// Пытаемся получить значение курсора и для последующей материализации
		if cursor, err = s.Storage.Find(*s.Context, filter); err == nil {
			defer cursor.Close(*s.Context)
			/// Пытаемся получить записи из курсора
			if err = cursor.All(*s.Context, &records); err == nil {
				/// Наполняем значениями RecordSet
				results.SetItems(records)
			} else {
				/// Если не удалось прочитать записи, то собираем ошибку
				err = &backbone_error.StoreError{
					Status:       error_constants.ERR_STORE_DECODE,
					StorageName:  s.Storage.Name(),
					DatabaseName: s.Storage.Database().Name(),
					Err:          err,
				}
			}
		} else {
			/// Если не удалось получить курсор, то собираем ошибку
			err = &backbone_error.StoreError{
				Status:       error_constants.ERR_STORE_READ,
				StorageName:  s.Storage.Name(),
				DatabaseName: s.Storage.Database().Name(),
				Err:          err,
			}
		}
	} else {
		/// Если не удалось вставить документы, то собираем ошибку
		err = &backbone_error.StoreError{
			Status:       error_constants.ERR_STORE_INSERT,
			StorageName:  s.Storage.Name(),
			DatabaseName: s.Storage.Database().Name(),
			Err:          err,
		}
	}

	return results, err
}

// InsertOne
// Метод вставки в коллекцию документа
func (s *MongoStore[DATA]) InsertOne(data DATA) (record.Record[primitive.ObjectID, DATA], error) {
	var err error
	var result *mongo.InsertOneResult
	var inserted record.Record[primitive.ObjectID, DATA]

	to_insert := &record.Record[primitive.ObjectID, DATA]{
		Identifier: primitive.NewObjectID(),
		Data:       data,
		CreatedAt:  &primitive.Timestamp{T: uint32(time.Now().Unix())},
	}

	if result, err = s.Storage.InsertOne(*s.Context, to_insert); err == nil {
		identifier_field := s.Filter["Identifier"]
		filter := bson.M{
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
