package initialize

import (
	"github.com/nas03/scholar-ai/backend/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedis() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Address,
		Password: global.Config.Redis.Password, // No password set
		DB:       global.Config.Redis.Database, // Use default DB
		Protocol: 2,                            // Connection protocol
	})
	global.Log.Info("Redis client established successfully",
		zap.String("address", global.Config.Redis.Address),
		zap.Int("database", global.Config.Redis.Database),
	)
	global.Redis = redisClient
}
