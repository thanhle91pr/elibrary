package model

type Combos struct {
	Books  string `json:"books" bson:"books"`
	Songs  string `json:"songs" bson:"songs"`
	Labels string `json:"labels" bson:"labels"`
}
