package model

type Label struct {
	ID          string   `json:"id" bson:"id"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	Books       string   `json:"books" bson:"books"`
	Songs       string   `json:"songs" bson:"songs"`
	Combos      []string `json:"combos" bson:"combos"`
}

type GetListLabelsReq struct {
	Page int64 `json:"p" query:"p"`
	Size int64 `json:"s" query:"s"`
}

type SetLabelsRequest struct {
	Labels string `json:"labels"`
	Books string `json:"books"`
	Songs string `json:"songs"`
}
