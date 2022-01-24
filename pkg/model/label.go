package model

type Label struct {
	ID          string  `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type GetListLabelsReq struct {
	Page int64 `json:"p" query:"p"`
	Size int64 `json:"s" query:"s"`
}
