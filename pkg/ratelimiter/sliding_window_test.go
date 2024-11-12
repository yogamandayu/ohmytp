package ratelimiter_test

import (
	"context"
	"fmt"
	"github.com/yogamandayu/ohmytp/internal/tests"
	"github.com/yogamandayu/ohmytp/pkg/ratelimiter"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestRateLimiterSlidingWindow(t *testing.T) {

	testSuite := tests.NewTestSuite()
	defer func() {
		t.Cleanup(testSuite.Clean)
	}()
	testSuite.LoadApp()

	t.Run("Positive case with 10 request and limit 10", func(t *testing.T) {
		fw := ratelimiter.NewSlidingWindow(testSuite.App.Log, testSuite.App.Redis).SetLimit(10)
		fw.SetRedisKey(fmt.Sprintf("rate_limit:sliding_window:%s", uuid.NewString()))
		var ok bool
		var err error
		for i := 0; i < 10; i++ {
			go func() {
				ok, err = fw.IsLimitReached(context.Background())
				require.NoError(t, err)
			}()
		}
		require.False(t, ok)
	})
	t.Run("Positive case with 10 request and limit 9", func(t *testing.T) {
		fw := ratelimiter.NewSlidingWindow(testSuite.App.Log, testSuite.App.Redis).SetLimit(9)
		fw.SetRedisKey(fmt.Sprintf("rate_limit:sliding_window:%s", uuid.NewString()))
		var ok bool
		var err error
		for i := 0; i < 10; i++ {
			go func() {
				ok, err = fw.IsLimitReached(context.Background())
				require.NoError(t, err)
			}()
		}
		require.True(t, ok)
	})
}
