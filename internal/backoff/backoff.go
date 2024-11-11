package backoff

import "time"

type Backoff struct {
	Jitter bool
	Min    time.Duration
	Max    time.Duration
}
