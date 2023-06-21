package mongo

import (
	database "github.com/small-entropy/go-backbone/database/abstract"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoBD struct {
	database.Database[*mongo.Client, *options.ClientOptions]
}
