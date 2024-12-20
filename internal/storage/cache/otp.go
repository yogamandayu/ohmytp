package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/yogamandayu/ohmytp/internal/domain/entity"

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

// SetOTP is to save request otp cache.
func (o *OTPCache) SetOTP(ctx context.Context, identifier string, otp entity.Otp, ttl time.Duration) {
	key := fmt.Sprintf("otp:%s", identifier)
	b, _ := json.Marshal(otp)
	o.redis.Set(ctx, key, string(b), ttl)
}

// GetOTP is to get request otp cache.
func (o *OTPCache) GetOTP(ctx context.Context, identifier string) (otp entity.Otp) {
	key := fmt.Sprintf("otp:%s", identifier)
	cmd := o.redis.Get(ctx, key)
	val := cmd.Val()
	_ = json.Unmarshal([]byte(val), &otp)
	return
}

// InvalidateOTP is to invalidate request otp cache.
func (o *OTPCache) InvalidateOTP(ctx context.Context, identifier string) {
	key := fmt.Sprintf("otp:%s", identifier)
	o.redis.Del(ctx, key)
}
