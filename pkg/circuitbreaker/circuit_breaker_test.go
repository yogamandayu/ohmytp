package circuitbreaker_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yogamandayu/ohmytp/pkg/circuitbreaker"
	"github.com/yogamandayu/ohmytp/tests"
	"testing"
	"time"
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
}
