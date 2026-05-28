package cache

import (
	"sync"
	"time"
)

type MemoryCache[K comparable, V any] struct {
	data map[K]*Entry[V]
	ttl  time.Duration
	mut  sync.RWMutex
}

func NewMemoryCache[K comparable, V any](ttl time.Duration, cleanupInterval time.Duration) *MemoryCache[K, V] {
	cache := MemoryCache[K, V]{
		data: make(map[K]*Entry[V]),
		ttl:  ttl,
		mut:  sync.RWMutex{},
	}

	go func() {
		for {
			<-time.After(cleanupInterval)
			cache.removeExpired()
		}
	}()

	return &cache
}

func (m *MemoryCache[K, V]) Set(key K, value V) {
	entry := NewEntry(value)

	m.mut.Lock()
	defer m.mut.Unlock()

	m.data[key] = &entry
}

func (m *MemoryCache[K, V]) Get(key K) (*V, bool) {
	m.mut.RLock()
	defer m.mut.RUnlock()

	if entry, ok := m.data[key]; ok {
		if entry.hasExpired(m.ttl) {
			return nil, false
		}

		return &entry.Value, true
	}

	return nil, false
}

func (m *MemoryCache[K, V]) removeExpired() {
	m.mut.Lock()
	defer m.mut.Unlock()

	for key, val := range m.data {
		if val.hasExpired(m.ttl) {
			delete(m.data, key)
		}
	}
}
