package meli

import (
	env "meli/environments"
	storage2 "meli/storage"
)

type Meli interface {
	GetClientID() string
	GetClientSecret() string
	GetUserID() string
	GetTenantID() string
	GetEnvironment() env.Environment
}

type meli struct {
	clientID     string
	clientSecret string
	userID       string
	tenantID     string
	environment  env.Environment
}

func (m meli) GetClientID() string {
	return m.clientID
}

func (m meli) GetClientSecret() string {
	return m.clientSecret
}

func (m meli) GetUserID() string {
	return m.userID
}

func (m meli) GetTenantID() string {
	return m.tenantID
}

func (m meli) GetEnvironment() env.Environment {
	return m.environment
}

// NewCustomClient retrieves a client with custom configuration for storage, site and environment
func NewCustomClient(clientID, clientSecret string, env env.Environment) Meli {
	return &meli{
		clientID:     clientID,
		clientSecret: clientSecret,
		environment:  env,
	}
}

// NewClient retrieves a client with pre-defined configuration
// storage used is in-memory and default site is Brazil
func NewClient(appID, clientSecret string) Meli {
	storage := storage2.NewInMemoryCache()
	return &meli{
		clientID:     appID,
		clientSecret: clientSecret,
		environment:  env.NewProductionEnv(env.BRASIL, env.NewConfiguration(storage)),
	}
}
