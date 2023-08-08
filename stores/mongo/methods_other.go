package mongo

import (
	error_constants "github.com/small-entropy/go-backbone/constants/error"
	backbone_error "github.com/small-entropy/go-backbone/error"
	"github.com/small-entropy/go-backbone/utils/convert"
)

// Метод получения количества записей с учетом фильтра
func (s *MongoStore[DATA]) GetCount(filter_map map[string]interface{}) (int64, error) {
	var count int64
	var err error

	filter := convert.MapToBsonM(filter_map)
	opts := GetCountOptions()

	if count, err = s.Storage.CountDocuments(*s.Context, filter, opts); err != nil {
		err = &backbone_error.StoreError{
			Status:       error_constants.ERR_STORE_COUNT,
			StorageName:  s.Storage.Name(),
			DatabaseName: s.Storage.Database().Name(),
			Err:          err,
		}
	}
	return count, err
}
