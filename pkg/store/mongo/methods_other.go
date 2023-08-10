package mongo

import (
	constants "github.com/small-entropy/go-backbone/internal/constants/error"

	errors "github.com/small-entropy/go-backbone/pkg/error"

	facade "github.com/small-entropy/go-backbone/third_party/facade/mongo"

	"github.com/small-entropy/go-backbone/tools/convert"
)

// GetCount
// Метод получения количества записей с учетом фильтра
func (s *Store[DATA]) GetCount(filter map[string]interface{}) (int64, error) {
	var count int64
	var err error

	currentFilter := convert.MapToBsonM(filter)

	opts := facade.GetCountOptions()

	if count, err = s.Storage.CountDocuments(*s.Context, currentFilter, opts); err != nil {
		err = &errors.StoreError{
			Status:       constants.ErrStoreCount,
			StorageName:  s.Storage.Name(),
			DatabaseName: s.Storage.Database().Name(),
			Err:          err,
		}
	}
	return count, err
}
