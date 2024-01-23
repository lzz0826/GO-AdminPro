package utils

import (
	"github.com/sony/sonyflake"
)

// 在 utils 包中创建 Sonyflake 实例，用于生成分布式 ID
var sf *sonyflake.Sonyflake

func init() {
	// 创建 Sonyflake 实例时，可以根据需要调整 sonyflake.Settings 的配置
	settings := sonyflake.Settings{}
	sf = sonyflake.NewSonyflake(settings)
}

// generateID 函数用于生成分布式 ID
func GenerateID() (int64, error) {
	// 生成分布式 ID
	id, err := sf.NextID()
	if err != nil {
		return 0, err
	}
	return int64(id), nil
}
