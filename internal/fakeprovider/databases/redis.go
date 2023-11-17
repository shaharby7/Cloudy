package databases

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func InitiateRedisClient() error {
	RedisClient = redis.NewClient(
		&redis.Options{Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_DOMAIN"), os.Getenv("REDIS_PORT"))},
	)
	return nil
}
