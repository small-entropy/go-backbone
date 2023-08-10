package controller

import (
	"github.com/small-entropy/go-backbone/pkg/datatypes/record"
	"github.com/small-entropy/go-backbone/pkg/datatypes/recordset"
	prov "github.com/small-entropy/go-backbone/pkg/provider/store"
	"github.com/small-entropy/go-backbone/pkg/store/abstract"

	"github.com/go-playground/validator/v10"
	facade "github.com/small-entropy/go-backbone/third_party/facade/mongo"
)

type Controller[CONN any, ID any, DATA any] struct {
	Validator *validator.Validate
	Fields    map[string]string
}

type IController[CONN any, ID any, DATA any] interface {
	InsertOne(data DATA, provider *prov.StoreProvider[CONN, ID, DATA]) (record.Record[ID, DATA], error)
	FindOne(identifier ID, deleted bool, other_filters *map[string]interface{}, provider *prov.StoreProvider[CONN, ID, DATA]) (record.Record[ID, DATA], error)
	FindOneByFilter(filter map[string]interface{}, provider *prov.StoreProvider[CONN, ID, DATA]) (record.Record[ID, DATA], error)
	Find(filter map[string]interface{}, page *abstract.Page, provider *prov.StoreProvider[CONN, ID, DATA]) (recordset.RecordSet[ID, DATA], error)
	UpdateOne(filter map[string]interface{}, update DATA, provider *prov.StoreProvider[CONN, ID, DATA]) (record.Record[ID, DATA], error)
	DeleteOne(filter map[string]interface{}, provider *prov.StoreProvider[CONN, ID, DATA]) (record.Record[ID, DATA], error)
	EraseOne(filter map[string]interface{}, provider *prov.StoreProvider[CONN, ID, DATA]) (record.Record[ID, DATA], error)
	GetTimeNow() facade.Timestamp
	GetValidator() *validator.Validate
	GetField(keys string) string
}
