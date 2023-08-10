package interfaces

import "github.com/small-entropy/go-backbone/pkg/type/entity/meta"

type ISet[E any] interface {
	Items() *[]E
	Meta() *meta.Meta
	Item(index int) (E, error)
	SetItems(value []E)
	SetMeta(value meta.Meta)
}
