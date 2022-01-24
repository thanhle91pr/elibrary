package model

type Songs struct {
	ID string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}
