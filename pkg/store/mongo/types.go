package mongo

import (
	"errors"

	"github.com/small-entropy/go-backbone/pkg/type/entity/meta"

	"github.com/small-entropy/go-backbone/pkg/store/abstract"
	facade "github.com/small-entropy/go-backbone/third_party/facade/mongo"
)

// Store[DATA any]
// Структура описывающая хранилище данных в MongoDB
type Store[DATA any] struct {
	abstract.Store[*facade.Collection, facade.ObjectID, DATA]
}

// Document[D any]
// Структура записи документа
type Document[D any] struct {
	Identifier facade.ObjectID   `bson:"_id" json:"Identifier,omitempty"`
	Data       D                 `bson:",inline"`
	CreatedAt  *facade.Timestamp `bson:"CreatedAt" json:"CreatedAt,omitempty"`
	UpdatedAt  *facade.Timestamp `bson:"UpdatedAt" json:"UpdatedAt,omitempty"`
	DeletedAt  *facade.Timestamp `bson:"DeletedAt" json:"DeletedAt,omitempty"`
}

// GetIdentifier
// Метод получения идентификатора документа
func (d *Document[D]) GetIdentifier() facade.ObjectID {
	return d.Identifier
}

// GetData
// Метод получения данных документа
func (d *Document[D]) GetData() D {
	return d.Data
}

// Created
// Метод получения даты и времени создания документа
func (d *Document[D]) Created() *facade.Timestamp {
	return d.CreatedAt
}

// Updated
// Метод получения даты и времени обновления документа
func (d *Document[D]) Updated() *facade.Timestamp {
	return d.UpdatedAt
}

// Deleted
// Метод получения даты и времени удаления докумета
func (d *Document[D]) Deleted() *facade.Timestamp {
	return d.UpdatedAt
}

// DocumentSet
// Структура последовательности документов MongoDB
type DocumentSet[D any] struct {
	items []Document[D]
	meta  meta.Meta
}

// Метод получения данных последовательности
func (ds *DocumentSet[D]) Items() *[]Document[D] {
	return &ds.items
}

// Meta
// Метод получения метаданных
func (ds *DocumentSet[D]) Meta() *meta.Meta {
	return &ds.meta
}

// Item
// Метод получения элемента по индексу
func (ds *DocumentSet[D]) Item(index int) (Document[D], error) {
	var err error
	var item Document[D]
	max := len(ds.items) - 1
	if max < index {
		err = errors.New("index to large")
	} else {
		item = ds.items[index]
	}
	return item, err
}

// SetItems
// Метод наполнения DocumentSet
func (ds *DocumentSet[D]) SetItems(value []Document[D]) {
	ds.items = value
}

// SetMeta
// Метод наполнения Meta
func (ds *DocumentSet[D]) SetMeta(value meta.Meta) {
	ds.meta = value
}
