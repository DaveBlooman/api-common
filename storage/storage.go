package storage

// Storage interface
type Storage interface {
	Get(key string) (string, *Error)
	Set(key string, data string) *Error
}
