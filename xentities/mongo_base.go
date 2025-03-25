package xentities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoBaseId struct {
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

type MongoBaseRecord struct {
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
	CreatedBy string    `json:"created_by,omitempty" bson:"created_by,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	UpdatedBy string    `json:"updated_by,omitempty" bson:"updated_by,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	DeletedBy string    `json:"deleted_by,omitempty" bson:"deleted_by,omitempty"`
}
