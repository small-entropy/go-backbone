package abstract

import (
	"context"

	"github.com/small-entropy/go-backbone/datatypes/record"
	"github.com/small-entropy/go-backbone/datatypes/recordset"
)

type Store[STORAGE any, ID any, DATA any] struct {
	Storage STORAGE
	Context *context.Context
	Filter  map[string]string
}

type Page struct {
	Limit int64
	Skip  int64
}

type IStore[ID any, DATA any] interface {
	// OTHER
	GetCount(filter map[string]interface{}) (int64, error)
	// CRUD
	InsertOne(data DATA) (record.Record[ID, DATA], error)
	FindOne(filter map[string]interface{}) (record.Record[ID, DATA], error)
	FindAll(page Page, filter map[string]interface{}) (recordset.RecordSet[ID, DATA], error)
	DeleteOne(filter map[string]interface{}) (record.Record[ID, DATA], error)
	UpdateOne(filter map[string]interface{}, update map[string]interface{}) (record.Record[ID, DATA], error)
}
