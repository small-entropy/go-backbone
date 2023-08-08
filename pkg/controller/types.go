package controller

import (
	"github.com/small-entropy/go-backbone/datatypes/record"
	"github.com/small-entropy/go-backbone/datatypes/recordset"
	store_provider "github.com/small-entropy/go-backbone/providers/store"
	"github.com/small-entropy/go-backbone/stores/abstract"

	"github.com/go-playground/validator/v10"
	mongo_facade "github.com/small-entropy/go-backbone/facade/mongo"
)

type Controller[CONN any, ID any, DATA any] struct {
	Validator *validator.Validate
	Fields    map[string]string
}

type IController[CONN any, ID any, DATA any] interface {
	InsertOne(data DATA, provider *store_provider.StoreProvider[CONN, ID, DATA]) (record.Record[ID, DATA], error)
	FindOne(identifier ID, deleted bool, other_filters *map[string]interface{}, provider *store_provider.StoreProvider[CONN, ID, DATA]) (record.Record[ID, DATA], error)
	FindOneByFilter(filter map[string]interface{}, provider *store_provider.StoreProvider[CONN, ID, DATA]) (record.Record[ID, DATA], error)
	Find(filter map[string]interface{}, page *abstract.Page, provider *store_provider.StoreProvider[CONN, ID, DATA]) (recordset.RecordSet[ID, DATA], error)
	UpdateOne(filter map[string]interface{}, update DATA, provider *store_provider.StoreProvider[CONN, ID, DATA]) (record.Record[ID, DATA], error)
	DeleteOne(filter map[string]interface{}, provider *store_provider.StoreProvider[CONN, ID, DATA]) (record.Record[ID, DATA], error)
	EraseOne(filter map[string]interface{}, provider *store_provider.StoreProvider[CONN, ID, DATA]) (record.Record[ID, DATA], error)
	GetTimeNow() mongo_facade.Timestamp
	GetValidator() *validator.Validate
	GetField(keys string) string
}
