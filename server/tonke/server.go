package tonke

import (
	"AdminPro/common/redis"
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
