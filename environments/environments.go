package environments

// Storage will be used to store request tokens
type Storage interface {
	Set(name, value string) error
	Get(name string) interface{}
	Has(name string) bool
	Remove(name string) error
}

// Configuration determine the resources that will be used
type Configuration interface {
	GetStorage() Storage
}

type configuration struct {
	storage Storage
}

func (c configuration) GetStorage() Storage {
	return c.storage
}

func NewConfiguration(storage Storage) Configuration {
	return &configuration{
		storage: storage,
	}
}
