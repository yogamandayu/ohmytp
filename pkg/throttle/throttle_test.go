package throttle_test

import (
	"context"
	"testing"
	"time"

	"github.com/yogamandayu/ohmytp/internal/tests"
	"github.com/yogamandayu/ohmytp/pkg/throttle"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestThrottle(t *testing.T) {

	testSuite := tests.NewTestSuite()
	defer func() {
		t.Cleanup(testSuite.Clean)
	}()
	testSuite.LoadApp()

	t.Run("Try to run process 3 times to test first threshold (3-5-7)", func(t *testing.T) {
		th := throttle.NewThrottle(testSuite.App.Redis, "example_process", uuid.NewString()).SetThresholds([]throttle.Threshold{
			{
				MaxAttempt:      3,
				WaitingDuration: 30 * time.Minute,
			},
			{
				MaxAttempt:      5,
				WaitingDuration: 60 * time.Minute,
			},
			{
				MaxAttempt:      7,
				WaitingDuration: 24 * 60 * time.Minute,
			},
		})
		for i := 0; i < 3; i++ {
			ok, err := th.IsAllowed(context.Background())
			require.NoError(t, err)
			assert.True(t, ok)
		}
	})

	t.Run("Try to run process 4 times to test first threshold (3-5-7)", func(t *testing.T) {
		th := throttle.NewThrottle(testSuite.App.Redis, "example_process", uuid.NewString()).SetThresholds([]throttle.Threshold{
			{
				MaxAttempt:      3,
				WaitingDuration: 30 * time.Minute,
			},
			{
				MaxAttempt:      5,
				WaitingDuration: 60 * time.Minute,
			},
			{
				MaxAttempt:      7,
				WaitingDuration: 24 * 60 * time.Minute,
			},
		})
		for i := 0; i < 4; i++ {
			ok, err := th.IsAllowed(context.Background())
			require.NoError(t, err)
			if i == 3 {
				assert.False(t, ok)
			}
		}
	})

	t.Run("Try to run process 8 times to test second threshold (3-5-7)", func(t *testing.T) {
		th := throttle.NewThrottle(testSuite.App.Redis, "example_process", uuid.NewString()).SetThresholds([]throttle.Threshold{
			{
				MaxAttempt:      3,
				WaitingDuration: 1 * time.Second,
			},
			{
				MaxAttempt:      5,
				WaitingDuration: 60 * time.Minute,
			},
			{
				MaxAttempt:      7,
				WaitingDuration: 24 * 60 * time.Minute,
			},
		})
		for i := 0; i < 8; i++ {
			ok, err := th.IsAllowed(context.Background())
			if i == 2 {
				time.Sleep(1 * time.Second)
			}
			require.NoError(t, err)
			assert.True(t, ok)
		}
	})

	t.Run("Try to run process 9 times to test second threshold (3-5-7)", func(t *testing.T) {
		th := throttle.NewThrottle(testSuite.App.Redis, "example_process", uuid.NewString()).SetThresholds([]throttle.Threshold{
			{
				MaxAttempt:      3,
				WaitingDuration: 1 * time.Second,
			},
			{
				MaxAttempt:      5,
				WaitingDuration: 60 * time.Minute,
			},
			{
				MaxAttempt:      7,
				WaitingDuration: 24 * 60 * time.Minute,
			},
		})
		for i := 0; i < 8; i++ {
			ok, err := th.IsAllowed(context.Background())
			if i == 2 {
				time.Sleep(1 * time.Second)
			}
			require.NoError(t, err)
			if i == 8 {
				assert.False(t, ok)
			}
		}
	})

	t.Run("Try to run process 15 times to test third threshold (3-5-7)", func(t *testing.T) {
		th := throttle.NewThrottle(testSuite.App.Redis, "example_process", uuid.NewString()).SetThresholds([]throttle.Threshold{
			{
				MaxAttempt:      3,
				WaitingDuration: 1 * time.Second,
			},
			{
				MaxAttempt:      5,
				WaitingDuration: 1 * time.Second,
			},
			{
				MaxAttempt:      7,
				WaitingDuration: 24 * 60 * time.Minute,
			},
		})
		for i := 0; i < 15; i++ {
			ok, err := th.IsAllowed(context.Background())
			if i == 2 || i == 7 {
				time.Sleep(1 * time.Second)
			}
			require.NoError(t, err)
			assert.True(t, ok)
		}
	})

	t.Run("Try to run process 9 times to test third threshold (3-5-7)", func(t *testing.T) {
		th := throttle.NewThrottle(testSuite.App.Redis, "example_process", uuid.NewString()).SetThresholds([]throttle.Threshold{
			{
				MaxAttempt:      3,
				WaitingDuration: 1 * time.Second,
			},
			{
				MaxAttempt:      5,
				WaitingDuration: 60 * time.Minute,
			},
			{
				MaxAttempt:      7,
				WaitingDuration: 24 * 60 * time.Minute,
			},
		})
		for i := 0; i < 15; i++ {
			ok, err := th.IsAllowed(context.Background())
			if i == 2 || i == 7 {
				time.Sleep(1 * time.Second)
			}
			require.NoError(t, err)
			if i == 15 {
				assert.False(t, ok)
			}
		}
	})

}
