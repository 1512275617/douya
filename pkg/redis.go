package pkg

import (
	"context"
	"douya/admin/config"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	redisClient *redis.Client
)

func RedisInit() func() {
	redisConfig := config.DefaultConfig.RedisCfg

	client := redis.NewClient(&redis.Options{
		Addr:         redisConfig.Addr,
		Password:     redisConfig.Password,
		DB:           redisConfig.DB,
		WriteTimeout: time.Duration(time.Duration(redisConfig.WriteTimeout).Seconds()),
		ReadTimeout:  time.Duration(time.Duration(redisConfig.ReadTimeout).Seconds()),
		DialTimeout:  time.Duration(time.Duration(redisConfig.DialTimeout).Seconds()),
		PoolSize:     4049,
	})

	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		panic(err)
	}

	redisClient = client

	return func() {
		err := client.Close()
		if err != nil {
			return
		}
		return
	}
}

func GetRedisClient(ctx context.Context) *redis.Client {
	return redisClient.WithContext(ctx)
}
