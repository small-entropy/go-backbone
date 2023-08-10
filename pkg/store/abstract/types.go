package abstract

import (
	"context"

	"github.com/small-entropy/go-backbone/pkg/type/interfaces"
)

// Store
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

type IStore[ID any, DATA any, DATETIME any, ENTITY any] interface {
	GetCount(filter map[string]interface{}) (int64, error)                                                                   // Получить количество записей по фильтру
	InsertOne(data DATA) (interfaces.ISet[ENTITY], error)                                                                                     // Вставить одну запись в хранилище
	InsertMany(data []DATA) (interfaces.ISet[ENTITY], error)                                                                 // Вставить несколько записей в хранилище
	FindOne(filter map[string]interface{}) (interfaces.ISchema[ID, DATA, DATETIME], error)                                   // Найти одну запись в хранилище по фильтру
	FindAll(page Page, filter map[string]interface{}) (interfaces.ISet[ENTITY], error)                                       // Найти все записи в хранилище по фильтру
	DeleteMany(filter map[string]interface{}) (interfaces.ISet[ENTITY], error)                                               // Удалить записи по фильтру
	DeleteOne(filter map[string]interface{}) (interfaces.ISchema[ID, DATA, DATETIME], error)                                 // Удалить одну запись по фильтру
	UpdateOne(filter map[string]interface{}, update map[string]interface{}) (interfaces.ISchema[ID, DATA, DATETIME], error)  // Обновить одну запись по фильтру
	UpdateMany(filter map[string]interface{}, update map[string]interface{}) (interfaces.ISchema[ID, DATA, DATETIME], error) // Обновить список записей по фильтру
}
