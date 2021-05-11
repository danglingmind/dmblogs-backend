package auth

import (
	"log"
	"strings"

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
		// redigo doesn't support username in the url, will remove the url for Heroku redis
		// for now
		// TODO: change the redis driver which supports user:password@url
		redisUrlWithoutUsername := strings.Replace(redis_url, "redistogo", "", 1)
		log.Printf("redis_url_without_username: %s", redisUrlWithoutUsername)
		redisClient, err = redis.DialURL(redisUrlWithoutUsername)
	}
	if err != nil {
		log.Fatalf("func: NewRedisDB error: %s redis_url: %s", err.Error(), redis_url)
		return nil, err
	}

	return &RedisService{
		Auth:   NewAuth(redisClient),
		Client: redisClient,
	}, nil
}
