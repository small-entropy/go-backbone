package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BsonM = bson.M

type BsonD = bson.D

type ObjectID = primitive.ObjectID

type InsertOneResult = mongo.InsertOneResult

type Timestamp = primitive.Timestamp

type CountOptions = options.CountOptions

type FindOptions = options.FindOptions

type Cursor = mongo.Cursor

type Collection = mongo.Collection

type Client = mongo.Client

type ClientOptions = options.ClientOptions

type IndexModel = mongo.IndexModel

type IndexOptions = options.IndexOptions

var ErrNoDocuments = mongo.ErrNoDocuments

var ErrNilDocument = mongo.ErrNilDocument

func NewClient(opts ...*options.ClientOptions) (*Client, error) {
	return mongo.NewClient(opts...)
}

func GetCountOptions() *CountOptions {
	return options.Count()
}

func GetFindOptions() *FindOptions {
	return options.Find()
}

func NewObjectID() primitive.ObjectID {
	return primitive.NewObjectID()
}

func GetClientOptions() *ClientOptions {
	return options.Client()
}

func GetIndexOptions() *IndexOptions {
	return options.Index()
}
