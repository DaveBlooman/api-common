package caching

import "time"

// Cache interface
type Cache interface {
	Get(key string) (string, error)
	Set(key string, data string, expiration time.Duration) (string, error)
}
