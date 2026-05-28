package cache

import "time"

type Entry[T any] struct {
	Value   T
	Updated time.Time
}

func NewEntry[T any](value T) Entry[T] {
	return Entry[T]{
		value,
		time.Now(),
	}
}

func (e *Entry[T]) hasExpired(ttl time.Duration) bool {
	return time.Since(e.Updated) >= ttl
}
