package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocMultImgs struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Images []Image            `json:"images,omitempty" bson:"images,omitempty"`
}
