package cache_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/th3shadowbroker/hymetric/internal/cache"
)

func TestCache(t *testing.T) {
	ttl, cleanupInterval := 2*time.Second, 4*time.Second
	testCache := cache.NewMemoryCache[string, string](ttl, cleanupInterval)

	t.Run("retrieve valid value", func(t *testing.T) {
		testCache.Set("foo", "bar")
		time.Sleep(ttl / 2)

		val, ok := testCache.Get("foo")
		assert.True(t, ok)
		assert.NotNil(t, val)
	})

	t.Run("retrieve expired value", func(t *testing.T) {
		testCache.Set("foo", "bar")
		time.Sleep(ttl + 500*time.Millisecond)

		val, ok := testCache.Get("foo")
		assert.False(t, ok)
		assert.Nil(t, val)
	})

	t.Run("retrieve value after cleanup interval", func(t *testing.T) {
		testCache.Set("foo", "bar")
		time.Sleep(cleanupInterval + 500*time.Millisecond)

		val, ok := testCache.Get("foo")
		assert.False(t, ok)
		assert.Nil(t, val)
	})
}
