package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/yogamandayu/ohmytp/consts"

	"github.com/yogamandayu/ohmytp/internal/app"
	"github.com/yogamandayu/ohmytp/internal/interfaces/rest/response"
	"github.com/yogamandayu/ohmytp/internal/requester"
	"github.com/yogamandayu/ohmytp/pkg/ratelimiter"
)

// RateLimiterMiddleware is a middleware to apply rate limiter to route handler.
type RateLimiterMiddleware struct {
	app *app.App

	fw *ratelimiter.FixedWindow

	strategy    consts.RateLimiterStrategy
	filterBy    consts.RateLimiterFilter
	handlerName string
}

// WithFixedWindow is to set rate limit with fixed window strategy.
func (rl *RateLimiterMiddleware) WithFixedWindow(limit int64, duration time.Duration) *RateLimiterMiddleware {
	rl.strategy = consts.FixedWindowStrategy
	rl.fw = ratelimiter.NewFixedWindow(rl.app.Log, rl.app.RedisAPI).SetLimit(limit).SetDuration(duration)
	return rl
}

// LimitByIPAddress is to rate limit by ip address.
func (rl *RateLimiterMiddleware) LimitByIPAddress() *RateLimiterMiddleware {
	rl.filterBy = consts.IPAddressFilter
	return rl
}

// RateLimit in init rate limit middleware.
func RateLimit(app *app.App, handlerName string) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		app:         app,
		strategy:    consts.FixedWindowStrategy,
		filterBy:    consts.IPAddressFilter,
		handlerName: handlerName,
	}
}

// Apply is to apply rate limit to handler. it's an end of constructing struct.
func (rl *RateLimiterMiddleware) Apply(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, _ := context.WithTimeout(r.Context(), 5*time.Second)
		rq := requester.NewRequester().SetMetadataFromREST(r)

		if rl.strategy == consts.FixedWindowStrategy {
			if rl.filterBy == consts.IPAddressFilter {
				rl.fw.SetRedisKey(fmt.Sprintf("rate_limit:fixed_window:%s:ip_address:%s", rl.handlerName, rq.Metadata.IPAddress))
			}
			ok, _ := rl.fw.IsLimitReached(ctx)
			if ok {
				response.NewHTTPFailedResponse("ERR101", errors.New("middleware.rate_limit.rate_limit_reached"), "Too Many Request").
					WithStatusCode(http.StatusTooManyRequests).
					AsJSON(w)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
