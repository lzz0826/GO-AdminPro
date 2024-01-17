package user

import (
	"AdminPro/dao/service/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetById(ctx *gin.Context) {

	id := ctx.Param("id")

	data, err := admin.GetById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
