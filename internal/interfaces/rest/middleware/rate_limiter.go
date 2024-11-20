package middleware

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/yogamandayu/ohmytp/internal/interfaces/rest/response"

	"github.com/yogamandayu/ohmytp/consts"

	"github.com/yogamandayu/ohmytp/internal/app"
	"github.com/yogamandayu/ohmytp/internal/requester"
	"github.com/yogamandayu/ohmytp/pkg/ratelimiter"
)

// RateLimiterMiddleware is a middleware to apply rate limiter to route handler.
type RateLimiterMiddleware struct {
	app *app.App

	Limit    int64
	Duration time.Duration

	rateLimiter ratelimiter.Interface

	strategy    consts.RateLimiterStrategy
	filter      consts.RateLimiterFilter
	processName string
}

// WithFixedWindow is to set rate limit with fixed window strategy.
func (rl *RateLimiterMiddleware) WithFixedWindow(limit int64, duration time.Duration) *RateLimiterMiddleware {
	rl.strategy = consts.FixedWindowStrategy
	rl.Limit = limit
	rl.Duration = duration
	return rl
}

// LimitByIPAddress is to rate limit by ip address.
func (rl *RateLimiterMiddleware) LimitByIPAddress() *RateLimiterMiddleware {
	rl.filter = consts.IPAddressFilter
	return rl
}

// SetProcessName is to set process name to set as part of identifier.
func (rl *RateLimiterMiddleware) SetProcessName(processName string) *RateLimiterMiddleware {
	rl.processName = processName
	return rl
}

// NewRateLimit is a constructor. Default rate limit is with Fixed Window strategy and with IP address as filter.
func NewRateLimit(app *app.App) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		app:         app,
		Limit:       100,
		Duration:    1 * time.Minute,
		strategy:    consts.FixedWindowStrategy,
		filter:      consts.IPAddressFilter,
		processName: "default_process",
	}
}

// Apply is to apply rate limit to handler. it's an end of constructing struct.
func (rl *RateLimiterMiddleware) Apply(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, _ := context.WithTimeout(r.Context(), 5*time.Second)

		switch rl.strategy {
		case consts.FixedWindowStrategy:
			fixedWindow := ratelimiter.NewFixedWindow(rl.app.Log, rl.app.RedisAPI).SetLimit(rl.Limit).SetDuration(rl.Duration)
			fixedWindow.SetRedisKey(rl.GenerateRedisKey(r))
			rl.rateLimiter = fixedWindow
		default:
			fixedWindow := ratelimiter.NewFixedWindow(rl.app.Log, rl.app.RedisAPI).SetLimit(100).SetDuration(1 * time.Minute)
			fixedWindow.SetRedisKey(rl.GenerateRedisKey(r))
			rl.rateLimiter = fixedWindow
		}

		ok, _ := rl.rateLimiter.IsLimitReached(ctx)
		if ok {
			response.NewHTTPFailedResponse("ERR101", errors.New("middleware.rate_limit.rate_limit_reached"), "Too Many Request").
				WithStatusCode(http.StatusTooManyRequests).
				AsJSON(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// GenerateRedisKey is to generate redis key based on strategy, process name, and filter.
// default key is rate_limit:fixed_window:default_process:ip_address:127.0.0.1.
func (rl *RateLimiterMiddleware) GenerateRedisKey(r *http.Request) string {
	var key string

	key = rl.KeyBuilder(key, "rate_limit")

	switch rl.strategy {
	case consts.FixedWindowStrategy:
		key = rl.KeyBuilder(key, "fixed_window")
	default:
		key = rl.KeyBuilder(key, "fixed_window")
	}

	if rl.processName != "" {
		key = rl.KeyBuilder(key, rl.processName)
	} else {
		key = rl.KeyBuilder(key, "default_process")
	}

	switch rl.filter {
	case consts.IPAddressFilter:
		rq := requester.NewRequester().SetMetadataFromREST(r)
		key = rl.KeyBuilder(key, "ip_address", rq.Metadata.IPAddress)
	default:
		rq := requester.NewRequester().SetMetadataFromREST(r)
		key = rl.KeyBuilder(key, "ip_address", rq.Metadata.IPAddress)
	}

	return key
}

// KeyBuilder is rate limit redis key appender.
func (rl *RateLimiterMiddleware) KeyBuilder(key string, identifiers ...string) string {
	for i, identifier := range identifiers {
		if key != "" || i < len(identifiers)-1 {
			key += ":"
		}
		key += identifier
	}
	return key
}
