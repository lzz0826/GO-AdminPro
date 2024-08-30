package model

type Pagination struct {
	Size int    `json:"size" form:"size" uri:"size"`
	Page int    `json:"page" form:"page" uri:"page"`
	Sort string `json:"sort" form:"sort" uri:"sort"`
}
