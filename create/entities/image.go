package entities

type Image struct {
	Data string `json:"data,omitempty" bson:"data,omitempty"`
	Alt string	`json:"alt,omitempty" bson:"alt,omitempty"`
}
