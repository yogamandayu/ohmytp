package circuitbreaker_test

import (
	"context"
	"testing"
	"time"

	"github.com/yogamandayu/ohmytp/consts"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yogamandayu/ohmytp/pkg/circuitbreaker"
	"github.com/yogamandayu/ohmytp/tests"
)

func TestCircuitBreaker(t *testing.T) {
	testSuite := tests.NewTestSuite()
	defer func() {
		t.Cleanup(testSuite.Clean)
	}()
	testSuite.LoadApp()

	t.Run("Request Success", func(t *testing.T) {
		cb := circuitbreaker.NewCircuitBreaker(testSuite.App.RedisAPI).
			SetPolicy(circuitbreaker.Policy{
				Timeframe:                 1 * time.Minute,
				CloseStateFailRate:        100,
				CloseStateFailDuration:    1 * time.Second,
				HalfOpenStateFailRate:     20,
				HalfOpenStateSuccessRate:  80,
				HalfOpenStateFailDuration: 1 * time.Second,
				OpenStateDuration:         30 * time.Second,
			}).SetRedisKey("test_success_1")
		for i := 0; i < 1000; i++ {
			ok, err := cb.IsAllowed(context.Background())
			require.NoError(t, err)
			assert.True(t, ok)
			err = cb.RecordSuccess(context.Background())
			require.NoError(t, err)
		}
	})
	t.Run("Request failed", func(t *testing.T) {
		cb := circuitbreaker.NewCircuitBreaker(testSuite.App.RedisAPI).
			SetPolicy(circuitbreaker.Policy{
				Timeframe:                 1 * time.Minute,
				CloseStateFailRate:        20,
				CloseStateFailDuration:    30 * time.Second,
				HalfOpenStateFailRate:     20,
				HalfOpenStateSuccessRate:  80,
				HalfOpenStateFailDuration: 30 * time.Second,
				OpenStateDuration:         30 * time.Second,
			}).SetRedisKey("test_fail_1")
		for i := 0; i < 1000; i++ {
			_, err := cb.IsAllowed(context.Background())
			require.NoError(t, err)
			if i >= 200 && i < 400 {
				err = cb.RecordError(context.Background())
				require.NoError(t, err)
			} else {
				err = cb.RecordSuccess(context.Background())
				require.NoError(t, err)
			}
		}
	})

	t.Run("Initial state should be CLOSED", func(t *testing.T) {
		cb := circuitbreaker.NewCircuitBreaker(testSuite.App.RedisAPI).SetRedisKey("test_cb").SetPolicy(circuitbreaker.Policy{
			Timeframe:                 10 * time.Second,
			CloseStateFailRate:        50.0,
			CloseStateFailDuration:    10 * time.Second,
			HalfOpenStateFailRate:     25.0,
			HalfOpenStateSuccessRate:  75.0,
			HalfOpenStateFailDuration: 5 * time.Second,
			OpenStateDuration:         15 * time.Second,
		})
		state, err := cb.DefineCurrentState(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, consts.CircuitBreakerStateClose.String(), state)
	})

	t.Run("Allow requests when state is CLOSED", func(t *testing.T) {
		cb := circuitbreaker.NewCircuitBreaker(testSuite.App.RedisAPI).SetRedisKey("test_cb").SetPolicy(circuitbreaker.Policy{
			Timeframe:                 10 * time.Second,
			CloseStateFailRate:        50.0,
			CloseStateFailDuration:    10 * time.Second,
			HalfOpenStateFailRate:     25.0,
			HalfOpenStateSuccessRate:  75.0,
			HalfOpenStateFailDuration: 5 * time.Second,
			OpenStateDuration:         15 * time.Second,
		})
		isAllowed, err := cb.IsAllowed(context.Background())
		assert.NoError(t, err)
		assert.True(t, isAllowed)
	})

	t.Run("Move to OPEN state after exceeding fail rate", func(t *testing.T) {
		cb := circuitbreaker.NewCircuitBreaker(testSuite.App.RedisAPI).SetRedisKey("test_cb").SetPolicy(circuitbreaker.Policy{
			Timeframe:                 10 * time.Second,
			CloseStateFailRate:        50.0,
			CloseStateFailDuration:    10 * time.Second,
			HalfOpenStateFailRate:     25.0,
			HalfOpenStateSuccessRate:  75.0,
			HalfOpenStateFailDuration: 5 * time.Second,
			OpenStateDuration:         15 * time.Second,
		})
		// Simulate errors
		for i := 0; i < 10; i++ {
			err := cb.RecordError(context.Background())
			assert.NoError(t, err)
		}

		state, err := cb.DefineCurrentState(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, consts.CircuitBreakerStateOpen.String(), state)
	})

	t.Run("Allow requests in HALF-OPEN state after open duration", func(t *testing.T) {
		cb := circuitbreaker.NewCircuitBreaker(testSuite.App.RedisAPI).SetRedisKey("test_cb").SetPolicy(circuitbreaker.Policy{
			Timeframe:                 10 * time.Second,
			CloseStateFailRate:        50.0,
			CloseStateFailDuration:    10 * time.Second,
			HalfOpenStateFailRate:     25.0,
			HalfOpenStateSuccessRate:  75.0,
			HalfOpenStateFailDuration: 5 * time.Second,
			OpenStateDuration:         15 * time.Second,
		})
		// Wait for OpenStateDuration
		time.Sleep(16 * time.Second)

		isAllowed, err := cb.IsAllowed(context.Background())
		assert.NoError(t, err)
		assert.True(t, isAllowed)

		state, err := cb.DefineCurrentState(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, consts.CircuitBreakerStateHalfOpen.String(), state)
	})

	t.Run("Transition back to CLOSED after successful requests in HALF-OPEN", func(t *testing.T) {
		cb := circuitbreaker.NewCircuitBreaker(testSuite.App.RedisAPI).SetRedisKey("test_cb").SetPolicy(circuitbreaker.Policy{
			Timeframe:                 10 * time.Second,
			CloseStateFailRate:        50.0,
			CloseStateFailDuration:    10 * time.Second,
			HalfOpenStateFailRate:     25.0,
			HalfOpenStateSuccessRate:  75.0,
			HalfOpenStateFailDuration: 5 * time.Second,
			OpenStateDuration:         15 * time.Second,
		})
		// Simulate successful requests
		for i := 0; i < 10; i++ {
			err := cb.RecordSuccess(context.Background())
			assert.NoError(t, err)
		}

		state, err := cb.DefineCurrentState(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, consts.CircuitBreakerStateClose.String(), state)
	})

	t.Run("Remain in HALF-OPEN after exceeding fail rate", func(t *testing.T) {
		cb := circuitbreaker.NewCircuitBreaker(testSuite.App.RedisAPI).SetRedisKey("test_cb").SetPolicy(circuitbreaker.Policy{
			Timeframe:                 10 * time.Second,
			CloseStateFailRate:        50.0,
			CloseStateFailDuration:    10 * time.Second,
			HalfOpenStateFailRate:     25.0,
			HalfOpenStateSuccessRate:  75.0,
			HalfOpenStateFailDuration: 5 * time.Second,
			OpenStateDuration:         15 * time.Second,
		})
		// Simulate errors in HALF-OPEN state
		cb.SetState(context.Background(), consts.CircuitBreakerStateHalfOpen)
		for i := 0; i < 5; i++ {
			err := cb.RecordError(context.Background())
			assert.NoError(t, err)
		}

		state, err := cb.DefineCurrentState(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, consts.CircuitBreakerStateOpen.String(), state)
	})
}
