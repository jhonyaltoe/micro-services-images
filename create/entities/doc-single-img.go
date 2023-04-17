package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocSingleImg struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Image Image              `json:"image,omitempty" bson:"image,omitempty" validate:"required"`
}
