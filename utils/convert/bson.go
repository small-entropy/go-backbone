package convert

import mongo_facade "github.com/small-entropy/go-backbone/facades/mongo"

// Метод конвертации Map в bson.M
func MapToBsonM(some_map map[string]interface{}) mongo_facade.BsonM {
	some_bson := mongo_facade.BsonM{}
	for k, v := range some_map {
		some_bson[k] = v
	}
	return some_bson
}
