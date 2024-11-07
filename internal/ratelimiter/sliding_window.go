package ratelimiter

import (
	"context"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

// SlidingWindow is rate limit strategy.
type SlidingWindow struct {
	RedisClient *redis.Client
	Log         *slog.Logger

	Duration time.Duration
	Limit    int64
	RedisKey string
}

var _ Interface = &SlidingWindow{}

// NewSlidingWindow is a constructor.
func NewSlidingWindow(log *slog.Logger, redis *redis.Client) *SlidingWindow {
	return &SlidingWindow{
		Duration:    1 * time.Minute,
		Limit:       100,
		RedisClient: redis,
		Log:         log,
	}
}

// SetLimit is to set rate limit capacity.
func (f *SlidingWindow) SetLimit(limit int64) *SlidingWindow {
	f.Limit = limit
	return f
}

// SetRedisKey is to set redis key to rate limit because this rate limit use redis.
func (f *SlidingWindow) SetRedisKey(key string) *SlidingWindow {
	f.RedisKey = key
	return f
}

// SetDuration is to set duration of each limit.
func (f *SlidingWindow) SetDuration(duration time.Duration) *SlidingWindow {
	f.Duration = duration
	return f
}

// ResetLimit is to reset limit of current redis key.
func (f *SlidingWindow) ResetLimit(ctx context.Context) {
	f.RedisClient.Del(ctx, f.RedisKey)
}

// IsLimitReached is to check is ip address already reach the rate limit.
func (f *SlidingWindow) IsLimitReached(ctx context.Context) (bool, error) {
	now := time.Now()
	nowStr := now.Format(time.RFC3339Nano)

	count, err := f.RedisClient.LLen(ctx, f.RedisKey).Result()
	if err != nil {
		return true, err
	}

	if count < f.Limit {
		_, err = f.RedisClient.LPush(ctx, f.RedisKey, nowStr).Result()
		if err != nil {
			return true, err
		}
		f.RedisClient.Expire(ctx, f.RedisKey, f.Duration)
		return false, nil
	}

	last, err := f.RedisClient.LIndex(ctx, f.RedisKey, -1).Result()
	if err != nil {
		return true, err
	}
	tLast, err := time.Parse(time.RFC3339Nano, last)
	if err != nil {
		return true, err
	}

	if now.Sub(tLast) > f.Duration {
		pipe := f.RedisClient.TxPipeline()
		pipe.LPush(ctx, f.RedisKey, nowStr)
		pipe.LTrim(ctx, f.RedisKey, 0, f.Limit-1)
		if _, err = pipe.Exec(ctx); err != nil {
			return true, err
		}
		return false, nil
	}
	return true, nil
}
