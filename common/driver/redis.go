package driver

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
	"strings"
	"time"
)

var (
	AdminRedisDb *RedisDb
	//不同的封装
	//UserRedisDb  *RedisDb
)

type RedisDb struct {
	aPool *redis.Client
}

func init() {
	client, err := initRedis()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	//这里赋予 RedisDb 可以个别处理不同的封装
	AdminRedisDb = &RedisDb{
		aPool: client,
	}
}

// 单节点 Redis 初始化
func initRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),     // Redis 地址
		Password: viper.GetString("redis.password"), // 如果没有设置密码，则留空
		DB:       viper.GetInt("redis.db"),          // 使用默认的 DB
	})

	// 测试连接
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}

// Redis 集群初始化
func initRedisCluster(host, auth string, poolSize int) (*redis.ClusterClient, error) {
	redisPool := redis.NewClusterClient(
		&redis.ClusterOptions{
			Addrs:              strings.Split(host, ","), // Redis 服务器地址列表，多个地址以逗号分隔，适用于 Redis 集群模式
			MaxRedirects:       0,                        // 最大重定向次数，适用于 Redis 集群，若为 0 则禁用重定向，默认值是 3
			ReadOnly:           false,                    // 是否只读模式，适用于 Redis 集群，若为 true 则命令只会发送到从节点
			RouteByLatency:     false,                    // 是否根据延迟路由命令到 Redis 节点，适用于 Redis 集群，延迟较低的节点会优先处理请求
			RouteRandomly:      false,                    // 是否随机路由命令到 Redis 节点，适用于 Redis 集群，若为 true 则随机选择一个节点处理请求
			Password:           auth,                     // Redis 连接密码，如果 Redis 服务器配置了密码验证，需要在这里填写
			MaxRetries:         2,                        // 最大重试次数，在命令失败时（如网络错误）自动重试，默认值为 3
			MinRetryBackoff:    8 * time.Millisecond,     // 最小重试等待时间，两次重试之间的最短等待时间
			MaxRetryBackoff:    512 * time.Millisecond,   // 最大重试等待时间，两次重试之间的最长等待时间
			DialTimeout:        5 * time.Second,          // 连接 Redis 服务器的超时时间，超过此时间未成功连接则返回错误
			ReadTimeout:        10 * time.Second,         // 读取 Redis 响应的超时时间，超过此时间未收到响应则返回错误
			WriteTimeout:       20 * time.Second,         // 发送命令到 Redis 的超时时间，超过此时间未成功发送则返回错误
			PoolSize:           poolSize,                 // 连接池大小，即可以同时连接到 Redis 的最大连接数
			MinIdleConns:       poolSize / 2,             // 连接池中保持的最小空闲连接数，设为 PoolSize 的一半以节省资源
			MaxConnAge:         6 * time.Minute,          // 连接的最大存活时间，超过此时间后连接会被自动关闭
			PoolTimeout:        30 * time.Second,         // 当连接池耗尽时，等待可用连接的超时时间
			IdleTimeout:        5 * time.Minute,          // 连接空闲的最大时间，超过此时间的空闲连接会被关闭
			IdleCheckFrequency: 1 * time.Minute,          // 检查空闲连接的频率，定期清理过期的空闲连接
		},
	)

	// 测试连接
	_, err := redisPool.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("host=%s |auth=%s |err=%v", host, auth, err)
		return nil, err
	}
	return redisPool, nil
}

// 获取指定 key 的字符串值
func (redisDb *RedisDb) GetKey(ctx context.Context, key string) (string, error) {
	value, err := redisDb.aPool.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return value, nil
}

// 获取指定 key 的字节值
func (redisDb *RedisDb) GetKeyBytes(ctx context.Context, key string) ([]byte, error) {
	return redisDb.aPool.Get(ctx, key).Bytes()
}

// 设置指定 key 的值且不设置过期时间
func (redisDb *RedisDb) SetNotExpireKV(ctx context.Context, key, value string) error {
	err := redisDb.aPool.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// 设置指定 key 的值，并设置过期时间
func (redisDb *RedisDb) SetExpireKV(ctx context.Context, key, value string, expire time.Duration) error {
	err := redisDb.aPool.Set(ctx, key, value, expire).Err()
	if err != nil {
		return err
	}
	return nil
}

// 获取指定 key 的剩余过期时间
func (redisDb *RedisDb) GetTTL(ctx context.Context, key string) time.Duration {
	return redisDb.aPool.TTL(ctx, key).Val()
}

// 为指定 key 设置过期时间
func (redisDb *RedisDb) Expire(ctx context.Context, key string, expire time.Duration) error {
	return redisDb.aPool.Expire(ctx, key, expire).Err()
}

// 为指定 key 设置过期时间（功能与 Expire 相同）
func (redisDb *RedisDb) SetExpireKey(ctx context.Context, key string, expire time.Duration) error {
	err := redisDb.aPool.Expire(ctx, key, expire).Err()
	if err != nil {
		return err
	}
	return nil
}

// 设置指定 key 的值，如果 key 不存在则设置成功
func (redisDb *RedisDb) SetNX(ctx context.Context, key string, value string, expire time.Duration) (bool, error) {
	flag, err := redisDb.aPool.SetNX(ctx, key, value, expire).Result()
	if err != nil {
		return false, err
	}
	return flag, nil
}

// 删除指定的 key
func (redisDb *RedisDb) DelKey(ctx context.Context, key string) error {
	return redisDb.aPool.Del(ctx, key).Err()
}

// 检查指定的 key 是否存在
func (redisDb *RedisDb) KeyExist(ctx context.Context, key string) (bool, error) {
	count, err := redisDb.aPool.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// 扫描指定模式匹配的 key
func (redisDb *RedisDb) Keys(ctx context.Context, pattern string) ([]string, error) {
	return redisDb.aPool.Keys(ctx, pattern).Result()
}

// 设置 hash 中的一个字段的值
func (redisDb *RedisDb) HSet(ctx context.Context, key, field string, value interface{}) error {
	return redisDb.aPool.HSet(ctx, key, field, value).Err()
}

// 批量设置 hash 中的多个字段的值
func (redisDb *RedisDb) HMSet(ctx context.Context, key string, fields map[string]interface{}) error {
	if len(fields) < 1 {
		return nil
	}

	err := redisDb.aPool.HMSet(ctx, key, fields).Err()
	if err != nil {
		return err
	}

	return nil
}

// 获取 hash 中指定字段的值
func (redisDb *RedisDb) HGet(ctx context.Context, key, field string) (string, error) {
	return redisDb.aPool.HGet(ctx, key, field).Result()
}

// 获取 hash 中所有字段的名称
func (redisDb *RedisDb) HKeys(ctx context.Context, key string) ([]string, error) {
	return redisDb.aPool.HKeys(ctx, key).Result()
}

// 批量获取 hash 中多个字段的值
func (redisDb *RedisDb) HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	res, err := redisDb.aPool.HMGet(ctx, key, fields...).Result()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 获取 hash 中所有字段和值
func (redisDb *RedisDb) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	res, err := redisDb.aPool.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 扫描 hash 中的字段
func (redisDb *RedisDb) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return redisDb.aPool.HScan(ctx, key, cursor, match, count).Result()
}

// 扫描集合中的成员
func (redisDb *RedisDb) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return redisDb.aPool.SScan(ctx, key, cursor, match, count).Result()
}

// 获取 hash 中字段的数量
func (redisDb *RedisDb) HLen(ctx context.Context, key string) (int, error) {
	res, err := redisDb.aPool.HLen(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return int(res), nil
}

// 删除 hash 中的一个或多个字段
func (redisDb *RedisDb) HDel(ctx context.Context, key string, fields ...string) error {
	err := redisDb.aPool.HDel(ctx, key, fields...).Err()
	if err != nil {
		return err
	}

	return nil
}

// 在列表的右侧推入一个或多个值
func (redisDb *RedisDb) RPush(ctx context.Context, key string, values ...interface{}) error {
	return redisDb.aPool.RPush(ctx, key, values...).Err()
}

// 在列表的左侧推入一个或多个值
func (redisDb *RedisDb) LPush(ctx context.Context, key string, values ...interface{}) error {
	return redisDb.aPool.LPush(ctx, key, values...).Err()
}

// 发布消息到频道
func (redisDb *RedisDb) Publish(ctx context.Context, channel string, values interface{}) error {
	return redisDb.aPool.Publish(ctx, channel, values).Err()
}

// 获取列表的长度
func (redisDb *RedisDb) LLen(ctx context.Context, key string) (int64, error) {
	return redisDb.aPool.LLen(ctx, key).Result()
}

// 获取列表中指定范围的元素
func (redisDb *RedisDb) LRange(ctx context.Context, key string, start, end int64) ([]string, error) {
	return redisDb.aPool.LRange(ctx, key, start, end).Result()
}

// 设置列表中指定索引的值
func (redisDb *RedisDb) LSet(ctx context.Context, key string, index int64, value interface{}) error {
	return redisDb.aPool.LSet(ctx, key, index, value).Err()
}

// 移除列表中指定数量的值
func (redisDb *RedisDb) LRem(ctx context.Context, key string, count int64, value interface{}) error {
	return redisDb.aPool.LRem(ctx, key, count, value).Err()
}

// 向有序集合添加成员及分数
func (redisDb *RedisDb) ZAdd(ctx context.Context, key, member string, score float64) error {
	z := redis.Z{
		Score:  score,
		Member: member,
	}
	_, err := redisDb.aPool.ZAdd(ctx, key, &z).Result()
	if err != nil {
		return err
	}

	return nil
}

// 统计有序集合中指定分数范围的成员数量
func (redisDb *RedisDb) ZCount(ctx context.Context, key, min, max string) (int64, error) {
	count, err := redisDb.aPool.ZCount(ctx, key, min, max).Result()
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 获取有序集合的成员数量
func (redisDb *RedisDb) ZCARD(ctx context.Context, key string) (int64, error) {
	count, err := redisDb.aPool.ZCard(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 获取有序集合中指定范围的成员
func (redisDb *RedisDb) ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	arr, err := redisDb.aPool.ZRange(ctx, key, start, stop).Result()
	if err != nil {
		return []string{}, err
	}

	return arr, nil
}
