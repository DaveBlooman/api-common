package caching

import (
	"time"

	"github.com/DaveBlooman/api-common/Godeps/_workspace/src/github.com/stretchr/testify/mock"
)

// MockCache setup
type MockCache struct {
	mock.Mock
}

// Get method
func (c *MockCache) Get(key string) (string, error) {
	args := c.Called(key)

	return args.String(0), args.Error(1)
}

// Set method
func (c *MockCache) Set(key string, data string, expiration time.Duration) (string, error) {
	args := c.Called(key, data, expiration)

	return args.String(0), args.Error(1)
}
