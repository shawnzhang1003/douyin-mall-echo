package redislock

import (
    "context"
    "fmt"
	"github.com/redis/go-redis/v9"
    "time"
)

// RedisLock 定义分布式锁结构体
type RedisLock struct {
    client    *redis.Client
    lockKey   string
    lockValue string
    expiration time.Duration
}

// NewRedisLock 创建一个新的 Redis 分布式锁实例
func NewRedisLock(client *redis.Client, lockKey string, expiration time.Duration) *RedisLock {
    return &RedisLock{
        client:    client,
        lockKey:   lockKey,
        lockValue: fmt.Sprintf("%d", time.Now().UnixNano()),
        expiration: expiration,
    }
}

// Acquire 尝试获取分布式锁
func (l *RedisLock) Acquire(ctx context.Context) (bool, error) {
    set, err := l.client.SetNX(ctx, l.lockKey, l.lockValue, l.expiration).Result()
    if err != nil {
        return false, err
    }
    return set, nil
}

// Release 释放分布式锁
func (l *RedisLock) Release(ctx context.Context) error {
    script := `
        if redis.call("GET", KEYS[1]) == ARGV[1] then
            return redis.call("DEL", KEYS[1])
        else
            return 0
        end
    `
    result, err := l.client.Eval(ctx, script, []string{l.lockKey}, l.lockValue).Result()
    if err != nil {
        return err
    }
    if result.(int64) == 0 {
        return fmt.Errorf("failed to release lock: lock not held or value mismatch")
    }
    return nil
}
