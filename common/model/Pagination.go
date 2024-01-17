package model

type Pagination struct {
	Limit int    `json:"limit" form:"limit" uri:"limit"`
	Page  int    `json:"page" form:"page" uri:"page"`
	Sort  string `json:"sort" form:"sort" uri:"sort"`
}
