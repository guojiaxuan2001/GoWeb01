package config

import (
	"github.com/go-redis/redis"
	"log"
)

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})

	_, err := RedisClient.Ping().Result()

	if err != nil {
		log.Fatal("Failed to connect to redis, get error ", err)
		return
	}

}
