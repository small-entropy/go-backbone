package recordset

import (
	"errors"

	"github.com/small-entropy/go-backbone/pkg/type/entity/meta"
	"github.com/small-entropy/go-backbone/pkg/type/entity/record"
)

type RecordSet struct {
	items []record.Record
	meta  meta.Meta
}

func (rs *RecordSet) Items() *[]record.Record {
	return &rs.items
}

func (rs *RecordSet) Meta() *meta.Meta {
	return &rs.meta
}

func (rs *RecordSet) Item(index int) (record.Record, error) {
	var err error
	var item record.Record
	max := len(rs.items) - 1
	if max < index {
		err = errors.New("index to large")
	} else {
		item = rs.items[index]
	}
	return item, err
}

func (rs *RecordSet) SetItems(value []record.Record) {
	rs.items = value
}

func (rs *RecordSet) SetMeta(value meta.Meta) {
	rs.meta = value
}
