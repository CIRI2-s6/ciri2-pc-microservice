package models

type Pagination struct {
	Skip   int                    `json:"skip,omitempty" validate:"required"`
	Limit  int                    `json:"limit,omitempty" validate:"required"`
	Query  string                 `json:"query,omitempty"`
	Filter map[string]interface{} `json:"filter,omitempty"`
	Sort   map[string]interface{} `json:"sort,omitempty"`
}
