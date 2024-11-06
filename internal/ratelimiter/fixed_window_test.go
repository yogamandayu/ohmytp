package ratelimiter_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yogamandayu/ohmytp/internal/ratelimiter"
	"github.com/yogamandayu/ohmytp/tests"
)

func TestRateLimiterFixedWindow(t *testing.T) {

	testSuite := tests.NewTestSuite()
	defer func() {
		t.Cleanup(testSuite.Clean)
	}()
	testSuite.LoadApp()

	t.Run("Positive case with one request", func(t *testing.T) {
		fw := ratelimiter.NewFixedWindow(testSuite.App.Log, testSuite.App.Redis).SetLimit(1)
		fw.SetRedisKey("rate_limit:fixed_window:ip_address:127.0.0.1")

		ok, err := fw.IsLimitReached(context.Background())
		require.NoError(t, err)
		assert.False(t, ok)
		fw.ResetLimit(context.Background())
	})

	t.Run("Negative case with 3 requests", func(t *testing.T) {
		fw := ratelimiter.NewFixedWindow(testSuite.App.Log, testSuite.App.Redis).SetLimit(2)
		fw.SetRedisKey("rate_limit:fixed_window:ip_address:127.0.0.1")

		for i := 0; i < 3; i++ {
			ok, err := fw.IsLimitReached(context.Background())
			require.NoError(t, err)
			if i == 2 {
				assert.True(t, ok)
			}
		}
		fw.ResetLimit(context.Background())
	})
}
