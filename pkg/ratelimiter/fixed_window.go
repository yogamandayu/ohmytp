package ratelimiter

import (
	"context"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

// FixedWindow is rate limit strategy.
type FixedWindow struct {
	RedisClient *redis.Client
	Log         *slog.Logger

	Duration time.Duration
	Limit    int64
	RedisKey string
}

var _ Interface = &FixedWindow{}

// NewFixedWindow is a constructor.
func NewFixedWindow(log *slog.Logger, redis *redis.Client) *FixedWindow {
	return &FixedWindow{
		Duration:    1 * time.Minute,
		Limit:       100,
		RedisClient: redis,
		Log:         log,
	}
}

// SetLimit is to set rate limit capacity.
func (f *FixedWindow) SetLimit(limit int64) *FixedWindow {
	f.Limit = limit
	return f
}

// SetRedisKey is to set redis key to rate limit because this rate limit use redis.
func (f *FixedWindow) SetRedisKey(key string) *FixedWindow {
	f.RedisKey = key
	return f
}

// SetDuration is to set duration of each limit.
func (f *FixedWindow) SetDuration(duration time.Duration) *FixedWindow {
	f.Duration = duration
	return f
}

// ResetLimit is to reset limit of current redis key.
func (f *FixedWindow) ResetLimit(ctx context.Context) {
	f.RedisClient.Del(ctx, f.RedisKey)
}

// IsLimitReached is to check is ip address already reach the rate limit.
func (f *FixedWindow) IsLimitReached(ctx context.Context) (bool, error) {
	count, err := f.RedisClient.Incr(ctx, f.RedisKey).Result()
	if err != nil {
		f.Log.Warn(err.Error())
		return true, err
	}
	if count == 1 {
		f.RedisClient.Expire(ctx, f.RedisKey, f.Duration)
	}
	return !(count <= f.Limit), nil
}
