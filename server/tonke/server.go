package tonke

import (
	"AdminPro/common/redis"
	"github.com/gin-gonic/gin"
	"time"
)

//TODO Redis

// SetTokenToRides Key token Value adminId
func SetTokenToRides(c *gin.Context, adminId, token string) error {
	redis := redis.AdminRedisDb
	// 设置过期时间为 30 分 待拉到配置文件
	expire := 30 * time.Minute
	err := redis.SetExpireKV(c, token, adminId, expire)
	if err != nil {
		return err
	}
	return nil
}

// GetTokenToRides Key token
func GetTokenToRides(c *gin.Context, token string) (string, error) {
	redis := redis.AdminRedisDb
	v, err := redis.GetKey(c, token)
	if err != nil {
		return "", nil
	}
	return v, nil
}

// RemoveTokenToRides Key token
func RemoveTokenToRides(c *gin.Context, token string) error {
	redis := redis.AdminRedisDb
	err := redis.DelKey(c, token)
	if err != nil {
		return err
	}
	return nil
}
