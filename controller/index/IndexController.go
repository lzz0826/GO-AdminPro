package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "index",
	})
}
