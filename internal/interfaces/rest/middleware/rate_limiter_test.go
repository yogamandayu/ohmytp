package middleware_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yogamandayu/ohmytp/internal/interfaces/rest/middleware"
	"github.com/yogamandayu/ohmytp/tests"
	"testing"
)

func TestRateLimiter(t *testing.T) {
	testSuite := tests.NewTestSuite()
	defer func() {
		t.Cleanup(testSuite.Clean)
	}()
	testSuite.LoadApp()

	t.Run("Test key builder", func(t *testing.T) {
		rateLimiter := middleware.NewRateLimit(testSuite.App)
		var key string
		key = rateLimiter.KeyBuilder("", "example")
		assert.Equal(t, "example", key)

		key = rateLimiter.KeyBuilder(key, "identifier_1")
		assert.Equal(t, "example:identifier_1", key)

		key = rateLimiter.KeyBuilder(key, "identifier_2", "identifier_3")
		assert.Equal(t, "example:identifier_1:identifier_2:identifier_3", key)

		key = rateLimiter.KeyBuilder(key, "identifier_4")
		assert.Equal(t, "example:identifier_1:identifier_2:identifier_3:identifier_4", key)

		key = rateLimiter.KeyBuilder(key, "identifier_5", "identifier_6")
		assert.Equal(t, "example:identifier_1:identifier_2:identifier_3:identifier_4:identifier_5:identifier_6", key)
	})

	t.Run("Test redis key generator", func(t *testing.T) {
		rateLimiter := middleware.NewRateLimit(testSuite.App)

		r := tests.FakeHTTPRequest()
		key := rateLimiter.GenerateRedisKey(r)
		assert.Equal(t, "rate_limit:fixed_window:default_process:ip_address:127.0.0.1", key)
		rateLimiter.SetProcessName("example")
		key = rateLimiter.GenerateRedisKey(r)
		assert.Equal(t, "rate_limit:fixed_window:example:ip_address:127.0.0.1", key)
	})
}
