package interfaces

import "github.com/small-entropy/go-backbone/pkg/store/abstract"

type IProvider[CONN any, ID any, DATA any, DATETIME any, ENTITY any] interface {
	GetStore() abstract.IStore[ID, DATA, DATETIME, ENTITY]
}
