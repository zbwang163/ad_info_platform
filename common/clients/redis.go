package clients

import "github.com/go-redis/redis"

var Redis *redis.Client

func InitRedis() {
	if Redis == nil {
		Redis = redis.NewClient(&redis.Options{
			Addr:     "8.136.119.124:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	}
}
