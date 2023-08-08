package mongo

import (
	constants "github.com/small-entropy/go-backbone/internal/constants/error"

	"github.com/small-entropy/go-backbone/pkg/datatypes/record"
	"github.com/small-entropy/go-backbone/pkg/datatypes/recordset"
	errors "github.com/small-entropy/go-backbone/pkg/error"

	facade "github.com/small-entropy/go-backbone/third_party/facade/mongo"

	"github.com/small-entropy/go-backbone/tools/convert"
)

// UpdateOne
// Метод обновления одного документа в коллекции
func (s *MongoStore[DATA]) UpdateOne(filter map[string]interface{}, update map[string]interface{}) (record.Record[facade.ObjectID, DATA], error) {
	var err error
	var result record.Record[facade.ObjectID, DATA]

	currentFilter := convert.MapToBsonM(filter)
	/// Собираем структуру для обновления
	toUpdate := convert.MapToBsonM(update)

	query := facade.BsonM{
		"$set": toUpdate,
	}

	if _, err = s.Storage.UpdateOne(*s.Context, currentFilter, query); err == nil {
		/// Находим обновленную запись
		err = s.Storage.FindOne(*s.Context, currentFilter).Decode(&result)
	} else {
		/// Если не удалось обновить запись - формируем ошибку
		err = &errors.StoreError{
			Status:       constants.ErrStoreUpdate,
			StorageName:  s.Storage.Name(),
			DatabaseName: s.Storage.Database().Name(),
			Err:          err,
		}
	}
	return result, err
}

// UpdateMany
// Метод обновления нескольких документов в коллекции
func (s *MongoStore[DATA]) UpdateMany(filter map[string]interface{}, update map[string]interface{}) (recordset.RecordSet[facade.ObjectID, DATA], error) {
	var err error
	var results recordset.RecordSet[facade.ObjectID, DATA]
	var records []record.Record[facade.ObjectID, DATA]
	// Собираем BsonM структуры по картам
	/// Формируем структуру для фильтра
	currentFilter := convert.MapToBsonM(filter)
	/// Формируем структуру для обновления
	toUpdate := convert.MapToBsonM(update)
	/// Собираем структуру запроса
	query := facade.BsonM{
		"$mul": toUpdate,
	}

	// Пытаемся обновить данные
	if _, err = s.Storage.UpdateMany(*s.Context, currentFilter, query); err == nil {
		/// Если данные успешно обновлены, то пытаемся получить все
		///  обновленные записи по текущему фильтру
		var cursor *facade.Cursor
		if cursor, err = s.Storage.Find(*s.Context, currentFilter); err == nil {
			/// Если удалось получить курсов, то пытаемся прочитать данные
			defer cursor.Close(*s.Context)
			if err = cursor.All(*s.Context, &records); err == nil {
				/// Если удалось прочитать данные, то задаем их как элементы для
				/// результирующего RecordSet
				results.SetItems(records)
			} else {
				/// Если не удалось начитать данные, то формируем ошибку
				err = &errors.StoreError{
					Status:       constants.ErrStoreDecode,
					StorageName:  s.Storage.Name(),
					DatabaseName: s.Storage.Database().Name(),
					Err:          err,
				}
			}
		}
	} else {
		// Если не удалось обновить данные, то формируем ошибку
		err = &errors.StoreError{
			Status:       constants.ErrStoreUpdate,
			StorageName:  s.Storage.Name(),
			DatabaseName: s.Storage.Database().Name(),
			Err:          err,
		}
	}

	return results, err
}
