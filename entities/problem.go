package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Problem struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Content   string             `bson:"content"`
	CreatedAt time.Time          `bson:"createdAt"`
	CreatedBy string             `bson:"createdBy"`

	UpdatedAt *time.Time `bson:"updatedAt,omitempty"`
	UpdatedBy *string    `bson:"updatedBy,omitempty"`

	Categories []string `bson:"categories,omitempty"`

	DeletedAt *time.Time `bson:"deletedAt,omitempty"`
	DeletedBy *string    `bson:"deletedBy,omitempty"`
}
