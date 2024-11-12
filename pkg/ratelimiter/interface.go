package ratelimiter

import "context"

type Interface interface {
	IsLimitReached(ctx context.Context) (bool, error)
}
