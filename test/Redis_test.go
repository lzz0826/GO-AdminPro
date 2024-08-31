package test

import (
	"AdminPro/common/driver"
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestSetExpireKV(t *testing.T) {
	key := "example_key"
	value := "test_value"
	expire := 10 * time.Second // 设置过期时间为 10 秒
	// 使用上下文
	ctx := context.Background()
	// 调用 SetExpireKV 方法
	err := driver.AdminRedisDb.SetExpireKV(ctx, key, value, expire)
	if err != nil {
		t.Fatalf("Error setting key with expiration: %v", err)
	}
}

func TestGetKey(t *testing.T) {

	ctx := context.Background()
	key := "example_key"
	value, err := driver.AdminRedisDb.GetKey(ctx, key)
	if err != nil {
		log.Fatalf("Error getting key: %v", err)
	}
	// 输出获取到的值
	if value == "" {
		fmt.Printf("Key %s does not exist\n", key)
	} else {
		fmt.Printf("Value for key %s: %s\n", key, value)
	}

}
