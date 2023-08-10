package mongo

import (
	"errors"

	"github.com/small-entropy/go-backbone/pkg/type/entity/meta"

	"github.com/small-entropy/go-backbone/pkg/store/abstract"
	facade "github.com/small-entropy/go-backbone/third_party/facade/mongo"
)

type Store[DATA any] struct {
	abstract.Store[*facade.Collection, facade.ObjectID, DATA]
}

type Document[D any] struct {
	Identifier facade.ObjectID   `bson:"_id" json:"Identifier,omitempty"`
	Data       D                 `bson:",inline"`
	CreatedAt  *facade.Timestamp `bson:"CreatedAt" json:"CreatedAt,omitempty"`
	UpdatedAt  *facade.Timestamp `bson:"UpdatedAt" json:"UpdatedAt,omitempty"`
	DeletedAt  *facade.Timestamp `bson:"DeletedAt" json:"DeletedAt,omitempty"`
}

func (d *Document[D]) GetIdentifier() facade.ObjectID {
	return d.Identifier
}

func (d *Document[D]) GetData() D {
	return d.Data
}

func (d *Document[D]) Created() *facade.Timestamp {
	return d.CreatedAt
}

func (d *Document[D]) Updated() *facade.Timestamp {
	return d.UpdatedAt
}

func (d *Document[D]) Deleted() *facade.Timestamp {
	return d.UpdatedAt
}

type DocumentSet[D any] struct {
	items []Document[D]
	meta  meta.Meta
}

func (ds *DocumentSet[D]) Items() *[]Document[D] {
	return &ds.items
}

func (ds *DocumentSet[D]) Meta() *meta.Meta {
	return &ds.meta
}

func (ds *DocumentSet[D]) Item(index int) (Document[D], error) {
	var err error
	var item Document[D]
	max := len(ds.items) - 1
	if max < index {
		err = errors.New("index to large")
	} else {
		item = ds.items[index]
	}
	return item, err
}

func (ds *DocumentSet[D]) SetItems(value []Document[D]) {
	ds.items = value
}

func (ds *DocumentSet[D]) SetMeta(value meta.Meta) {
	ds.meta = value
}
