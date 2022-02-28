package meli

import (
	env "meli/environments"
)

type Meli interface {
	GetClientID() string
	GetClientSecret() string
	GetUserID() string
	GetTenantID() string
	GetEnvironment() env.EnvironmentConfig
}

type meli struct {
	clientID     string
	clientSecret string
	userID       string
	tenantID     string
	environment  env.EnvironmentConfig
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

func (m meli) GetEnvironment() env.EnvironmentConfig {
	return m.environment
}

func NewClient(clientID, clientSecret string, env env.EnvironmentConfig) Meli {
	return &meli{
		clientID:     clientID,
		clientSecret: clientSecret,
		environment:  env,
	}
}
