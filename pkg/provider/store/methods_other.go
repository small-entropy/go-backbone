package store

import "github.com/small-entropy/go-backbone/pkg/store/abstract"

func (p *Provider[CONN, ID, DATA, DATETIME, ENTITY]) GetStore() abstract.IStore[ID, DATA, DATETIME, ENTITY] {
	return p.Store
}
