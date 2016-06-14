package storage

import "github.com/stretchr/testify/mock"

// MockStorage mock
type MockStorage struct {
	mock.Mock
}

// Get method
func (s *MockStorage) Get(key string) (string, *Error) {
	args := s.Called(key)
	return args.String(0), args.Get(1).(*Error)
}

// Set method
func (s *MockStorage) Set(key string, data string) *Error {
	args := s.Called(key, data)
	return args.Get(0).(*Error)
}
