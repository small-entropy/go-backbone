package mongo

import (
	mongo_facade "github.com/small-entropy/go-backbone/facades/mongo"
	abstract_store "github.com/small-entropy/go-backbone/stores/abstract"
)

type MongoStore[DATA any] struct {
	abstract_store.Store[*mongo_facade.Collection, mongo_facade.ObjectID, DATA]
}
