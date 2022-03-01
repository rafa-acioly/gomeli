package storage

// Storage will be used to store request tokens
type Storage interface {
	Set(name, value string) error
	Get(name string) (string, error)
	Remove(name string) error
}
