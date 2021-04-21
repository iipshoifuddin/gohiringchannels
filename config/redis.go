package config

import (
	"os"

	"github.com/go-redis/redis"
)

type redisClient struct {
	c *redis.Client
}

var (
	client = &redisClient{}
)

//GetClient get the redis client
func SetupRedis() *redis.Client {
	ENVSetup()
	c := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	if err := c.Ping().Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}
	return c
}
