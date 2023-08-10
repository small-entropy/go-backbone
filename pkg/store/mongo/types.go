package mongo

import (
	"github.com/small-entropy/go-backbone/pkg/store/abstract"
	facade "github.com/small-entropy/go-backbone/third_party/facade/mongo"
)

type MongoStore[DATA any] struct {
	abstract.Store[*facade.Collection, facade.ObjectID, DATA]
}
