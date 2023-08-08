package recordset

import "github.com/small-entropy/go-backbone/pkg/datatypes/record"

// Get Item mehtod
func (rs *RecordSet[ID, DATA]) GetItems() []record.Record[ID, DATA] {
	return rs.Items
}

func (rs *RecordSet[ID, DATA]) SetItems(items []record.Record[ID, DATA]) {
	rs.Items = items
	rs.Meta.Count = len(items)
}

func (rs *RecordSet[ID, DATA]) GetCount() int {
	count := len(rs.Items)
	return count
}
