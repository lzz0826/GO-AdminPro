package utils

import (
	"AdminPro/common/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GeneratePaginationFromRequest(c *gin.Context) (pagination model.Pagination) {
	if err := c.ShouldBind(&pagination); err != nil {
		fmt.Printf("参数绑定错误:%s\n", err)
	}
	// 校验参数
	if pagination.Size < 0 {
		pagination.Size = 2
	}
	if pagination.Page < 1 {
		pagination.Page = 1
	}
	if len(pagination.Sort) == 0 {
		pagination.Sort = "create_time desc"
	}
	if pagination.Sort == "desc" {
		pagination.Sort = "create_time desc"
	}

	if pagination.Sort == "asc" {
		pagination.Sort = "create_time asc"
	}
	return
}
