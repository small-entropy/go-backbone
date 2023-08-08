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

func GetCountOptions() *CountOptions {
	return options.Count()
}

func GetFindOptions() *FindOptions {
	return options.Find()
}

func NewObjectID() primitive.ObjectID {
	return primitive.NewObjectID()
}
