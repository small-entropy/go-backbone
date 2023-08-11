package mongo

import (
	"github.com/small-entropy/go-backbone/pkg/controller/abstract"
	facade "github.com/small-entropy/go-backbone/third_party/facade/mongo"
)

type Controller[DATA any, ENTITY any] struct {
	abstract.Controller[*facade.Client, facade.ObjectID, DATA, facade.Timestamp, ENTITY]
}
