package environments

import "meli/storage"

// Configuration determine the resources that will be used
type Configuration interface {
	GetStorage() storage.Storage
}

type configuration struct {
	storage storage.Storage
}

func (c configuration) GetStorage() storage.Storage {
	return c.storage
}

func NewConfiguration(storage storage.Storage) Configuration {
	return &configuration{
		storage: storage,
	}
}
