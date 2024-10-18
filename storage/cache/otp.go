package cache

import "github.com/redis/go-redis/v9"

// OTPCache is a struct for caching OTP.
type OTPCache struct {
	redis *redis.Client
}

// NewOTPCache is a constructor.
func NewOTPCache(redis *redis.Client) *OTPCache {
	return &OTPCache{
		redis: redis,
	}
}
