package store

import (
	"context"

	"github.com/small-entropy/go-backbone/stores/abstract"
)

type StoreProvider[CONN any, ID any, DATA any] struct {
	Store   abstract.IStore[ID, DATA]
	Context *context.Context
}
