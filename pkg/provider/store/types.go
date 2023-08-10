package store

import (
	"context"

	"github.com/small-entropy/go-backbone/pkg/store/abstract"
)

type Provider[CONN any, ID any, DATA any, DATETIME any, ENTITY any] struct {
	Store   abstract.IStore[ID, DATA, DATETIME, ENTITY]
	Context *context.Context
}
