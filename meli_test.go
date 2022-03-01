package meli

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"meli/environments"
	"meli/storage"
	"testing"
)

type MeliTestSuite struct {
	config environments.Configuration
	cache  storage.Storage
	env    environments.Environment
	suite.Suite
}

func (test MeliTestSuite) SetupTest() {
	test.cache = storage.NewInMemoryCache()
	test.config = environments.NewConfiguration(test.cache)
	test.env = environments.NewSandboxEnv(environments.BRASIL, test.config)
}

func (test MeliTestSuite) CanCreateAClientInstance() {
	cli := NewClient("client-id", "client-secret", test.env)

	assert.Equal(test.T(), cli.GetClientID(), "client-id")
	assert.Equal(test.T(), cli.GetClientSecret(), "client-secret")
}

func TestMeliTestSuite(t *testing.T) {
	suite.Run(t, new(MeliTestSuite))
}
