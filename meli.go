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
	SetTenantID(key string)
	GetEnvironment() env.Environment
}

type meli struct {
	clientID     string
	clientSecret string
	userID       string
	tenantID     string
	environment  env.Environment
}

// GetClientID return the client id
func (m meli) GetClientID() string {
	return m.clientID
}

// GetClientSecret return the client secret
func (m meli) GetClientSecret() string {
	return m.clientSecret
}

// GetUserID return the user id
func (m meli) GetUserID() string {
	return m.userID
}

// GetTenantID return the tenant id
func (m meli) GetTenantID() string {
	return m.tenantID
}

// SetTenantID define the tenant id, this value is mainly used
// when you need to handle multiple users on your app
func (m *meli) SetTenantID(key string) {
	m.tenantID = key
}

// GetEnvironment return the current environment
func (m meli) GetEnvironment() env.Environment {
	return m.environment
}

// NewCustomClient retrieves a client with custom configuration for storage, site and environment
func NewCustomClient(clientID, clientSecret, tenantID string, env env.Environment) Meli {
	return &meli{
		clientID:     clientID,
		clientSecret: clientSecret,
		tenantID:     tenantID,
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
		tenantID:     "golang_sdk",
		environment:  env.NewProductionEnv(env.BRASIL, env.NewConfiguration(storage)),
	}
}
