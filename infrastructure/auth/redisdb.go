package auth

import (
	"github.com/gomodule/redigo/redis"
)

type RedisService struct {
	Auth   AuthInterface
	Client redis.Conn
}

func NewRedisDB(redis_url string) (*RedisService, error) {
	var redisClient redis.Conn
	var err error
	if redis_url == "" {
		redisClient, err = redis.Dial("tcp", ":6379")
	} else {
		redisClient, err = redis.DialURL(redis_url)
	}
	if err != nil {
		return nil, err
	}

	return &RedisService{
		Auth:   NewAuth(redisClient),
		Client: redisClient,
	}, nil
}
