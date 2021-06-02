package auth

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

type RedisService struct {
	Auth   AuthInterface
	Client redis.Conn
}

func NewRedisDB(redisHost string) (*RedisService, error) {
	redisClient, err := redis.Dial("tcp", redisHost+":6379")

	if err != nil {
		log.Fatalf("func: NewRedisDB error: %s redisHost: %s", err.Error(), redisHost)
		return nil, err
	}

	return &RedisService{
		Auth:   NewAuth(redisClient),
		Client: redisClient,
	}, nil
}
