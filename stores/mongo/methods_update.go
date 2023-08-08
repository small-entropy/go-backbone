package mongo

import (
	error_constants "github.com/small-entropy/go-backbone/constants/error"
	"github.com/small-entropy/go-backbone/datatypes/record"
	"github.com/small-entropy/go-backbone/datatypes/recordset"
	backbone_error "github.com/small-entropy/go-backbone/error"
	mongo_facade "github.com/small-entropy/go-backbone/facades/mongo"
	"github.com/small-entropy/go-backbone/utils/convert"
)

// UpdateOne
// Метод обновления одного документа в коллекции
func (s *MongoStore[DATA]) UpdateOne(filter map[string]interface{}, update map[string]interface{}) (record.Record[mongo_facade.ObjectID, DATA], error) {
	var err error
	var result record.Record[mongo_facade.ObjectID, DATA]

	filter_bson := convert.MapToBsonM(filter)
	/// Собираем структуру для обновления
	update_bson := convert.MapToBsonM(update)

	update_query := mongo_facade.BsonM{
		"$set": update_bson,
	}

	if _, err = s.Storage.UpdateOne(*s.Context, filter_bson, update_query); err == nil {
		/// Находим обновленную запись
		err = s.Storage.FindOne(*s.Context, filter_bson).Decode(&result)
	} else {
		/// Если не удалось обновить запись - формируем ошибку
		err = &backbone_error.StoreError{
			Status:       error_constants.ERR_STORE_UPDATE,
			StorageName:  s.Storage.Name(),
			DatabaseName: s.Storage.Database().Name(),
			Err:          err,
		}
	}
	return result, err
}

// UpdateMany
// Метод обновления нескольких документов в коллекции
func (s *MongoStore[DATA]) UpdateMany(filter map[string]interface{}, update map[string]interface{}) (recordset.RecordSet[mongo_facade.ObjectID, DATA], error) {
	var err error
	var results recordset.RecordSet[mongo_facade.ObjectID, DATA]
	var records []record.Record[mongo_facade.ObjectID, DATA]
	// Собираем BsonM структуры по картам
	/// Формируем структуру для фильтра
	filter_bson := convert.MapToBsonM(filter)
	/// Формируем структуру для обновления
	update_bson := convert.MapToBsonM(update)
	/// Собираем структуру запроса
	update_query := mongo_facade.BsonM{
		"$mul": update_bson,
	}

	// Пытаемся обновить данные
	if _, err = s.Storage.UpdateMany(*s.Context, filter_bson, update_query); err == nil {
		/// Если данные успешно обновлены, то пытаемся получить все
		///  обновленные записи по текущему фильтру
		var cursor *mongo_facade.Cursor
		if cursor, err = s.Storage.Find(*s.Context, filter_bson); err == nil {
			/// Если удалось получить курсов, то пытаемся прочитать данные
			defer cursor.Close(*s.Context)
			if err = cursor.All(*s.Context, &records); err == nil {
				/// Если удалось прочитать данные, то задаем их как элементы для
				/// результирующего RecordSet
				results.SetItems(records)
			} else {
				/// Если не удалось начитать данные, то формируем ошибку
				err = &backbone_error.StoreError{
					Status:       error_constants.ERR_STORE_DECODE,
					StorageName:  s.Storage.Name(),
					DatabaseName: s.Storage.Database().Name(),
					Err:          err,
				}
			}
		}
	} else {
		// Если не удалось обновить данные, то формируем ошибку
		err = &backbone_error.StoreError{
			Status:       error_constants.ERR_STORE_UPDATE,
			StorageName:  s.Storage.Name(),
			DatabaseName: s.Storage.Database().Name(),
			Err:          err,
		}
	}

	return results, err
}
