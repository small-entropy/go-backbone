package store

import "github.com/small-entropy/go-backbone/stores/abstract"

func (sp *StoreProvider[CONN, ID, DATA]) GetStore() abstract.IStore[ID, DATA] {
	return sp.Store
}
