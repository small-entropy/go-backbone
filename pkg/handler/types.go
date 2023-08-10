package handler

import (
	"context"

	"github.com/small-entropy/go-backbone/pkg/controller"
	"github.com/small-entropy/go-backbone/pkg/datatypes/record"
	"github.com/small-entropy/go-backbone/pkg/datatypes/recordset"
	provider "github.com/small-entropy/go-backbone/pkg/provider/store"
	"github.com/small-entropy/go-backbone/pkg/response/jsend"
)

type Fields struct {
	Response map[string]string
	Params   map[string]string
	Filter   map[string]string
	Messages map[string]string
}

type Callbacks[CONN any, ID any, DATA any, DTO any] struct {
	GetIdentifierFromString func(value string) (ID, error)
	GetProvider             func(ctx context.Context, storageName string) (provider.StoreProvider[CONN, ID, DATA], error)
	Fill                    func(dto DTO) (DATA, error)
}

type Settings[CONN any, ID any, DATA any] struct {
	Controller  controller.IController[CONN, ID, DATA]
	Connection  CONN
	StorageName string
	Fields      Fields
}

type CustomResponses[CONN any, ID any, DATA any, DTO any] struct {
	Item map[string]func(context *context.Context, handler *Handler[CONN, ID, DATA, DTO], rec *record.Record[ID, DATA], args ...interface{}) *jsend.Response
	List map[string]func(context *context.Context, handler *Handler[CONN, ID, DATA, DTO], recs *recordset.RecordSet[ID, DATA], args ...interface{}) *jsend.Response
}

type Responses[CONN any, ID any, DATA any, DTO any] struct {
	Name            string
	CreateOne       func(context *context.Context, handler *Handler[CONN, ID, DATA, DTO], rec *record.Record[ID, DATA], args ...interface{}) *jsend.Response
	GetOne          func(context *context.Context, handler *Handler[CONN, ID, DATA, DTO], rec *record.Record[ID, DATA], args ...interface{}) *jsend.Response
	List            func(context *context.Context, handler *Handler[CONN, ID, DATA, DTO], recs *recordset.RecordSet[ID, DATA], args ...interface{}) *jsend.Response
	UpdateOne       func(context *context.Context, handler *Handler[CONN, ID, DATA, DTO], rec *record.Record[ID, DATA], args ...interface{}) *jsend.Response
	DeleteOne       func(context *context.Context, handler *Handler[CONN, ID, DATA, DTO], rec *record.Record[ID, DATA], args ...interface{}) *jsend.Response
	EraseOne        func(context *context.Context, handler *Handler[CONN, ID, DATA, DTO], rec *record.Record[ID, DATA], args ...interface{}) *jsend.Response
	CustomResponses CustomResponses[CONN, ID, DATA, DTO]
}

type IAdapter[CONN any, ID any, DATA any, DTO any] interface {
	One(rec *record.Record[ID, DATA], args ...interface{}) interface{}
	List(recs *recordset.RecordSet[ID, DATA], args ...interface{}) []interface{}
}

type Handler[CONN any, ID any, DATA any, DTO any] struct {
	Settings  Settings[CONN, ID, DATA]
	Callbacks Callbacks[CONN, ID, DATA, DTO]
	Adapter   IAdapter[CONN, ID, DATA, DTO]
	Responses Responses[CONN, ID, DATA, DTO]
}
