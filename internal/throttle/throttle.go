package throttle

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Throttle struct {
	Redis      *redis.Client
	Process    string
	Identifier string

	Thresholds []Threshold
}

type Threshold struct {
	MaxAttempt      uint8
	WaitingDuration time.Duration
}

// Attempt is a redis key and value. format : throttle:{{process}}:{{identifier}} = {"total_attempt":1,"last_attempt_at":time.Now()}
type Attempt struct {
	CurrentAttempt uint8     `json:"current_attempt"`
	WaitUntil      time.Time `json:"wait_until"`
}

func NewThrottle(redis *redis.Client, processName string, identifier string) *Throttle {
	return &Throttle{
		Redis:      redis,
		Process:    processName,
		Identifier: identifier,
	}
}

func (t *Throttle) Reset(ctx context.Context) error {
	key := fmt.Sprintf("throttle:%s:%s", t.Process, t.Identifier)
	return t.Redis.Del(ctx, key).Err()
}

func (t *Throttle) IsAllowed(ctx context.Context) (bool, error) {

	if len(t.Thresholds) == 0 {
		return true, nil
	}

	var attempt Attempt
	now := time.Now()

	key := fmt.Sprintf("throttle:%s:%s", t.Process, t.Identifier)
	val, _ := t.Redis.Get(ctx, key).Result()
	if val == "" {
		attempt = Attempt{
			CurrentAttempt: 0,
		}
	}

	if val != "" {
		err := json.Unmarshal([]byte(val), &attempt)
		if err != nil {
			return false, err
		}
	}

	if now.Before(attempt.WaitUntil) {
		return false, nil
	}

	attempt.CurrentAttempt++
	if attempt.CurrentAttempt == t.ThresholdTotalAttemptByCurrentAttempt(attempt.CurrentAttempt) {
		threshold := t.ThresholdByCurrentAttempt(attempt.CurrentAttempt)
		attempt.WaitUntil = time.Now().Add(threshold.WaitingDuration)
	}

	b, _ := json.Marshal(attempt)
	err := t.Redis.Set(ctx, key, string(b), 24*time.Hour).Err()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (t *Throttle) ThresholdByCurrentAttempt(currentAttempt uint8) Threshold {
	var totalAttempt uint8
	for _, threshold := range t.Thresholds {
		totalAttempt += threshold.MaxAttempt
		if totalAttempt >= currentAttempt {
			return threshold
		}
	}
	return Threshold{}
}

func (t *Throttle) ThresholdTotalAttemptByCurrentAttempt(currentAttempt uint8) uint8 {
	var totalAttempt uint8
	for _, threshold := range t.Thresholds {
		totalAttempt += threshold.MaxAttempt
		if totalAttempt >= currentAttempt {
			return totalAttempt
		}
	}
	return totalAttempt
}

func (t *Throttle) SetThresholds(thresholds []Threshold) *Throttle {
	t.Thresholds = thresholds
	return t
}
