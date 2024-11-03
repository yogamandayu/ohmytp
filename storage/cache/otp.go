package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

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

// SetRequestOTP is to save request otp cache.
func (o *OTPCache) SetRequestOTP(ctx context.Context, requestID, otp string, ttl time.Duration) {
	key := fmt.Sprintf("otp:%s", requestID)
	o.redis.Set(ctx, key, otp, ttl)
}

// GetRequestOTP is to get request otp cache.
func (o *OTPCache) GetRequestOTP(ctx context.Context, requestID string) (val any) {
	key := fmt.Sprintf("otp:%s", requestID)
	cmd := o.redis.Get(ctx, key)
	val = cmd.Val()
	return
}
