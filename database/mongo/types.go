package mongo

import (
	database "github.com/small-entropy/go-backbone/database/abstract"
	mongo_facade "github.com/small-entropy/go-backbone/facades/mongo"
)

type MongoBD struct {
	database.Database[*mongo_facade.Client, *mongo_facade.ClientOptions]
}
