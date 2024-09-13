package utils

import "github.com/gin-gonic/gin"

// GetGinContextKey 从上下文中获取某个Key的值
func GetGinContextKey(c *gin.Context, key string) interface{} {
	value, ok := c.Get(key)
	if !ok {
		return nil
	}
	return value
}

// SetGinContextKey 向上下文中设置一个Key-Value
func SetGinContextKey(c *gin.Context, key string, value interface{}) bool {
	if key != "" && value != nil {
		c.Set(key, value)
		return true
	}
	return false
}
