package circuitbreaker

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yogamandayu/ohmytp/consts"
)

type CircuitBreaker struct {
	RedisClient *redis.Client
	RedisKey    string
	State       consts.CircuitBreakerState
	Policy      Policy
}

type Policy struct {
	Timeframe time.Duration

	CloseStateFailRate     float64
	CloseStateFailDuration time.Duration

	HalfOpenStateFailRate     float64
	HalfOpenStateSuccessRate  float64
	HalfOpenStateFailDuration time.Duration

	OpenStateDuration time.Duration
}

func NewCircuitBreaker(redis *redis.Client) *CircuitBreaker {
	return &CircuitBreaker{
		RedisClient: redis,
	}
}

func (c *CircuitBreaker) SetPolicy(policy Policy) *CircuitBreaker {
	c.Policy = policy
	return c
}

func (c *CircuitBreaker) SetRedisKey(redisKey string) *CircuitBreaker {
	c.RedisKey = redisKey
	return c
}

func (c *CircuitBreaker) IsAllowed(ctx context.Context) (bool, error) {

	currentState, err := c.DefineCurrentState(ctx)
	if err != nil {
		return false, err
	}

	switch currentState {
	case consts.CircuitBreakerStateClose.String():
		return true, nil
	case consts.CircuitBreakerStateHalfOpen.String():
		return true, nil
	case consts.CircuitBreakerStateOpen.String():
		stateUpdatedAtKey := "circuit_breaker:" + c.RedisKey + ":state_updated_at"
		updatedAt, err := c.RedisClient.Get(ctx, stateUpdatedAtKey).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			return false, err
		}
		updatedAtT, err := time.Parse(time.RFC3339, updatedAt)
		if err != nil {
			return false, err
		}
		if updatedAtT.Add(c.Policy.OpenStateDuration).Before(time.Now()) {
			err = c.SetState(ctx, consts.CircuitBreakerStateHalfOpen)
			if err != nil {
				return false, err
			}
			return true, nil
		}

		return false, nil
	}

	return true, nil
}

func (c *CircuitBreaker) DefineCurrentState(ctx context.Context) (string, error) {
	stateKey := "circuit_breaker:" + c.RedisKey + ":state"
	currentState, err := c.RedisClient.Get(ctx, stateKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return "", err
	}

	if errors.Is(err, redis.Nil) {
		currentState, err = c.RedisClient.Set(ctx, stateKey, consts.CircuitBreakerStateClose.String(), 0).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			return "", err
		}
	}

	errorCountKey := "circuit_breaker:" + c.RedisKey + ":error_count"
	_, err = c.RedisClient.Get(ctx, errorCountKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return "", err
	}
	if errors.Is(err, redis.Nil) {
		return consts.CircuitBreakerStateClose.String(), nil
	}
	return currentState, nil
}

func (c *CircuitBreaker) RecordSuccess(ctx context.Context) error {
	stateKey := "circuit_breaker:" + c.RedisKey + ":state"
	errorCountKey := "circuit_breaker:" + c.RedisKey + ":error_count"
	successCountKey := "circuit_breaker:" + c.RedisKey + ":success_count"
	requestCountKey := "circuit_breaker:" + c.RedisKey + ":request_count"

	cmd := c.RedisClient.Incr(ctx, requestCountKey)
	requestCount, err := cmd.Result()
	if err != nil {
		return err
	}
	if requestCount == 1 {
		c.RedisClient.Expire(ctx, requestCountKey, c.Policy.Timeframe)
	}

	currentState, err := c.RedisClient.Get(ctx, stateKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}

	switch currentState {
	case consts.CircuitBreakerStateClose.String():
		c.RedisClient.Del(ctx, errorCountKey)
		return nil
	case consts.CircuitBreakerStateHalfOpen.String():
		cmd := c.RedisClient.Incr(ctx, successCountKey)
		count, err := cmd.Result()
		if err != nil {
			return err
		}
		if count == 1 {
			c.RedisClient.Expire(ctx, successCountKey, 0)
		}
		successRate := (float64(count) / float64(requestCount)) * 100
		if successRate >= c.Policy.HalfOpenStateSuccessRate {
			err = c.SetState(ctx, consts.CircuitBreakerStateClose)
			if err != nil {
				return err
			}
			c.RedisClient.Del(ctx, successCountKey)
			c.RedisClient.Del(ctx, errorCountKey)
		}
		return nil
	}

	return nil
}

func (c *CircuitBreaker) RecordError(ctx context.Context) error {
	stateKey := "circuit_breaker:" + c.RedisKey + ":state"
	errorCountKey := "circuit_breaker:" + c.RedisKey + ":error_count"
	successCountKey := "circuit_breaker:" + c.RedisKey + ":success_count"
	requestCountKey := "circuit_breaker:" + c.RedisKey + ":request_count"

	cmd := c.RedisClient.Incr(ctx, requestCountKey)
	requestCount, err := cmd.Result()
	if err != nil {
		return err
	}
	if requestCount == 1 {
		c.RedisClient.Expire(ctx, requestCountKey, c.Policy.Timeframe)
	}

	currentState, err := c.RedisClient.Get(ctx, stateKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}

	switch currentState {
	case consts.CircuitBreakerStateClose.String():
		cmd := c.RedisClient.Incr(ctx, errorCountKey)
		count, err := cmd.Result()
		if err != nil {
			return err
		}
		if count == 1 {
			c.RedisClient.Expire(ctx, errorCountKey, c.Policy.CloseStateFailDuration)
		}
		errRate := (float64(count) / float64(requestCount)) * 100
		if errRate >= c.Policy.CloseStateFailRate {
			err = c.SetState(ctx, consts.CircuitBreakerStateOpen)
			if err != nil {
				return nil
			}
			return nil
		}

	case consts.CircuitBreakerStateHalfOpen.String():
		cmd := c.RedisClient.Incr(ctx, errorCountKey)
		count, err := cmd.Result()
		if err != nil {
			return err
		}
		if count == 1 {
			c.RedisClient.Expire(ctx, errorCountKey, c.Policy.HalfOpenStateFailDuration)
		}

		errRate := (float64(count) / float64(requestCount)) * 100
		if errRate >= c.Policy.HalfOpenStateFailRate {
			err = c.SetState(ctx, consts.CircuitBreakerStateOpen)
			if err != nil {
				return nil
			}
			c.RedisClient.Del(ctx, errorCountKey)
			c.RedisClient.Del(ctx, successCountKey)
			return nil
		}
	}

	return nil
}

func (c *CircuitBreaker) SetState(ctx context.Context, state consts.CircuitBreakerState) error {
	stateKey := "circuit_breaker:" + c.RedisKey + ":state"
	err := c.RedisClient.Set(ctx, stateKey, state.String(), 0).Err()
	if err != nil {
		return err
	}
	stateUpdatedAtKey := "circuit_breaker:" + c.RedisKey + ":state_updated_at"
	err = c.RedisClient.Set(ctx, stateUpdatedAtKey, time.Now().Format(time.RFC3339), 0).Err()
	if err != nil {
		return err
	}
	return nil
}
