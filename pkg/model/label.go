package model

type Label struct {
	ID          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name" validate:"presence,min=1,max=100"`
	Description string `json:"description" bson:"description" validate:"presence,min=1,max=200"`
	Books       string `json:"books" bson:"books"`
	Songs       string `json:"songs" bson:"songs"`
}

type GetListLabelsReq struct {
	Page int64 `json:"p" query:"p"`
	Size int64 `json:"s" query:"s"`
}

type SetLabelsRequest struct {
	Label string `json:"label"`
	Book  string `json:"book"`
	Songs string `json:"song"`
}

type FindLabelsRequest struct {
	Books string `json:"book" query:"book"`
	Song  string `json:"song" query:"song"`
}
