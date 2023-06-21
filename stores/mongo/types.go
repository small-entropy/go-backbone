package mongo

import (
	abstract_store "github.com/small-entropy/go-backbone/stores/abstract"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStore[DATA any] struct {
	abstract_store.Store[*mongo.Collection, primitive.ObjectID, DATA]
}
