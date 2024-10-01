package cache

import "github.com/redis/go-redis/v9"

type OTPCache struct {
	redis *redis.Client
}
