package utils

import (
	"context"
	"time"

	"github.com/nas03/scholar-ai/backend/global"
	"github.com/redis/go-redis/v9"
)

type IRedisCache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, data any) error
	SetEx(ctx context.Context, key string, data any, exp time.Duration) error
}

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache() IRedisCache {
	return &RedisCache{
		client: global.Redis,
	}
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisCache) Set(ctx context.Context, key string, data any) error {
	return r.client.Set(ctx, key, data, 0).Err()
}

func (r *RedisCache) SetEx(ctx context.Context, key string, data any, exp time.Duration) error {
	return r.client.SetEx(ctx, key, data, exp).Err()
}
