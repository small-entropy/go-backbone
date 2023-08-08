package abstract

import (
	"context"

	"github.com/small-entropy/go-backbone/pkg/datatypes/record"
	"github.com/small-entropy/go-backbone/pkg/datatypes/recordset"
)

// Структура источника данных
type Store[STORAGE any, ID any, DATA any] struct {
	Storage STORAGE           //хранилище
	Context *context.Context  // контекст выполнения
	Filter  map[string]string // конфигурация фильтров
}

type Page struct {
	Limit int64 // максимальное количество записей, которое нужно вернуть
	Skip  int64 // количество записей, которое надо "выкинуть" из поискового запроса (от начала списка данных)
}

type IStore[ID any, DATA any] interface {
	GetCount(filter map[string]interface{}) (int64, error)                                                          // Получить количество записей по фильтру
	InsertOne(data DATA) (record.Record[ID, DATA], error)                                                           // Вставить одну запись в хранилище
	InsertMany(data []DATA) (recordset.RecordSet[ID, DATA], error)                                                  // Вставить несколько записей в хранилище
	FindOne(filter map[string]interface{}) (record.Record[ID, DATA], error)                                         // Найти одну запись в хранилище по фильтру
	FindAll(page Page, filter map[string]interface{}) (recordset.RecordSet[ID, DATA], error)                        // Найти все записи в хранилище по фильтру
	DeleteMany(filter map[string]interface{}) (recordset.RecordSet[ID, DATA], error)                                // Удалить записи по фильтру
	DeleteOne(filter map[string]interface{}) (record.Record[ID, DATA], error)                                       // Удалить одну запись по фильтру
	UpdateOne(filter map[string]interface{}, update map[string]interface{}) (record.Record[ID, DATA], error)        // Обновить одну запись по фильтру
	UpdateMany(filter map[string]interface{}, update map[string]interface{}) (recordset.RecordSet[ID, DATA], error) // Обновить список записей по фильтру
}
