package initialize

import (
	"github.com/nas03/scholar-ai/backend/global"
	"github.com/redis/go-redis/v9"
)

func InitRedis() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Address,
		Password: global.Config.Redis.Password, // No password set
		DB:       global.Config.Redis.Database, // Use default DB
		Protocol: 2,                            // Connection protocol
	})

	global.Redis = redisClient
}
