package convert

import facade "github.com/small-entropy/go-backbone/third_party/facade/mongo"

// MapToBsonM
// Метод конвертации Map в bson.M
func MapToBsonM(some_map map[string]interface{}) facade.BsonM {
	result := facade.BsonM{}

	for k, v := range some_map {
		result[k] = v
	}

	return result
}
