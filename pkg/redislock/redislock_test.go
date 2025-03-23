package redislock

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

func TestRedisLock(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	lockKey := "test_lock1"
	expiration := 10 * time.Second
	lock := NewRedisLock(client, lockKey, expiration)

	ctx := context.Background()
	// 尝试获取锁
	locked, err := lock.Acquire(ctx)
	if err != nil {
		t.Errorf("Failed to acquire lock: %v", err)
	}
	if !locked {
		t.Errorf("Failed to acquire lock: lock is already held")
	}

	// 模拟业务操作
	time.Sleep(5 * time.Second)

	// 释放锁
	err = lock.Release(ctx)
	if err != nil {
		t.Errorf("Failed to release lock: %v", err)
	}
}
