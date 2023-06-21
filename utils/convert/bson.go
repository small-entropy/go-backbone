package convert

import "go.mongodb.org/mongo-driver/bson"

// Метод конвертации Map в bson.M
func MapToBsonM(some_map map[string]interface{}) bson.M {
	some_bson := bson.M{}
	for k, v := range some_map {
		some_bson[k] = v
	}
	return some_bson
}
