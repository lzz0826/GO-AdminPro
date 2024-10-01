package tonke

import (
	"AdminPro/common/redis"
	"AdminPro/internal/glog"
	"github.com/gin-gonic/gin"
	"time"
)

// SetTokenToRides Key adminId Value token
func SetTokenToRides(c *gin.Context, adminId, token string) error {
	redis := redis.AdminRedisDb
	//TODO 设置过期时间为 30 分 待拉到配置文件
	expire := 30 * time.Minute
	err := redis.SetExpireKV(c, adminId, token, expire)
	if err != nil {
		return err
	}
	return nil
}

// GetTokenToRides Key adminId
func GetTokenToRides(c *gin.Context, adminId string) (string, error) {
	redis := redis.AdminRedisDb
	v, err := redis.GetKey(c, adminId)
	if err != nil {
		return "", nil
	}
	return v, nil
}

// RemoveTokenToRides Key adminId
func RemoveTokenToRides(c *gin.Context, adminId string) error {
	redis := redis.AdminRedisDb
	err := redis.DelKey(c, adminId)
	if err != nil {
		return err
	}
	return nil
}

// 驗證前端和登入後的Token是否一致
func VerifyOnlineToken(c *gin.Context, adminId, token string) bool {
	ridesToken, err := GetTokenToRides(c, adminId)
	if ridesToken == "" || err != nil || token != ridesToken {
		glog.Errorf("token驗證錯誤 緩存Token:%d 驗證Token:%d", ridesToken, token)
		return false
	}
	return true
}
