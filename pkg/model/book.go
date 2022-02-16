package model

type Books struct {
	ID          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name" validate:"presence,min=1,max=100"`
	Description string `json:"description" bson:"description" validate:"presence,min=1,max=200"`
	Labels      string `json:"labels" bson:"labels""`
}
