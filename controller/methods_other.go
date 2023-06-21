package controller

import (
	"time"

	"github.com/go-playground/validator/v10"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Метод получения текущей даты в формате Timestamp
func (c *Controller[CONN, ID, DATA]) GetTimeNow() primitive.Timestamp {
	now := primitive.Timestamp{T: uint32(time.Now().Unix())}
	return now
}

func (c *Controller[CONN, ID, DATA]) GetValidator() *validator.Validate {
	return c.Validator
}

func (c *Controller[CONN, ID, DATA]) GetField(key string) string {
	return c.Fields[key]
}
