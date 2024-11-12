package consts

type RateLimiterStrategy string

const (
	FixedWindowStrategy   RateLimiterStrategy = "FIXED_WINDOW"
	SlidingWindowStrategy RateLimiterStrategy = "SLIDING_WINDOW"
)

type RateLimiterFilter string

const (
	IPAddressFilter RateLimiterFilter = "IP_ADDRESS"
)
