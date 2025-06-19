package utils

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"strings"
)

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

// 取得GET POST 參數
func GetAllRequestParams(c *gin.Context) map[string]interface{} {
	params := make(map[string]interface{})

	// 1. Query 参数
	for key, values := range c.Request.URL.Query() {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}

	// 2. Form 表单参数（包括 x-www-form-urlencoded 和 multipart/form-data）
	_ = c.Request.ParseForm()
	for key, values := range c.Request.PostForm {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}

	// 3. JSON Body 参数
	if c.Request.Method == http.MethodPost &&
		strings.Contains(c.GetHeader("Content-Type"), "application/json") {

		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 复原 Body

		var jsonBody map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &jsonBody); err == nil {
			for key, val := range jsonBody {
				params[key] = val
			}
		}
	}
	return params
}

func GetIntParam(params map[string]interface{}, key string) int {
	if val, ok := params[key]; ok {
		switch v := val.(type) {
		case float64:
			return int(v)
		case string:
			if n, err := strconv.Atoi(v); err == nil {
				return n
			}
		}
	}
	return 0
}
