package service

import (
	"github.com/gin-gonic/gin"
)

// GetCurrentAdminId 從上下文中獲取管理員ID
func GetCurrentAdminId(c *gin.Context) string {
	adminIdInterface, ok := c.Get("adminId")
	if !ok {
		return ""
	}
	adminId, ok := adminIdInterface.(string)
	if !ok {
		return ""
	}
	return adminId
}
