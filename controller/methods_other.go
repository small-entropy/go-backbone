package controller

import (
	"time"

	"github.com/go-playground/validator/v10"

	mongo_facade "github.com/small-entropy/go-backbone/facades/mongo"
)

// GetTimeNow
// Метод получения текущей даты в формате Timestamp
func (c *Controller[CONN, ID, DATA]) GetTimeNow() mongo_facade.Timestamp {
	now := mongo_facade.Timestamp{T: uint32(time.Now().Unix())}
	return now
}

// GetValidator
func (c *Controller[CONN, ID, DATA]) GetValidator() *validator.Validate {
	return c.Validator
}

// GetField
func (c *Controller[CONN, ID, DATA]) GetField(key string) string {
	return c.Fields[key]
}
