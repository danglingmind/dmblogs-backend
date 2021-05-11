package auth

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

type RedisService struct {
	Auth   AuthInterface
	Client redis.Conn
}

func NewRedisDB(redis_url string) (*RedisService, error) {
	var redisClient redis.Conn
	var err error
	// if redis_url == "" {
	// 	redisClient, err = redis.Dial("tcp", ":6379")
	// } else {
	// 	redisClient, err = redis.DialURL(redis_url)
	// }
	log.Printf("func: NewRedisDB  redis_url: %s", redis_url)
	redisClient, err = redis.DialURL(redis_url, redis.DialPassword("487331a69ff8cd65996d9e5d6728218e"))
	if err != nil {
		log.Fatalf("func: NewRedisDB error: %s", err.Error())
		return nil, err
	}

	return &RedisService{
		Auth:   NewAuth(redisClient),
		Client: redisClient,
	}, nil
}
