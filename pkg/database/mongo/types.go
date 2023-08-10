package mongo

import (
	database "github.com/small-entropy/go-backbone/pkg/database/abstract"
	facade "github.com/small-entropy/go-backbone/third_party/facade/mongo"
)

type MongoBD struct {
	database.Database[*facade.Client, *facade.ClientOptions]
}
