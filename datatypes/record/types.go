package record

import "go.mongodb.org/mongo-driver/bson/primitive"

// TODO: отвязять от MongoDB модель Record
type Record[ID any, DATA any] struct {
	Identifier ID                   `bson:"_id" json:"Identifier,omitempty"`
	Data       DATA                 `bson:",inline"`
	CreatedAt  *primitive.Timestamp `bson:"CreatedAt" json:"CreatedAt,omitempty"`
	UpdatedAt  *primitive.Timestamp `bson:"UpdatedAt" json:"UpdatedAt,omitempty"`
	DeletedAt  *primitive.Timestamp `bson:"DeletedAt" json:"DeletedAt,omitempty"`
}
