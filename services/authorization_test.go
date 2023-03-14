package services_test

import (
	"math"
	"meli"
	"meli/environments"
	"meli/services"
	"meli/storage"
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestAuthorizationService struct {
	m       meli.Meli
	env     environments.Environment
	service services.AuthorizationService

	suite.Suite
}

func (test *TestAuthorizationService) SetupTest() {
	test.env = environments.NewSandboxEnv(
		environments.BRASIL,
		environments.NewConfiguration(storage.NewInMemoryCache()),
	)
	test.m = meli.NewCustomClient("1618031856055412", "ItuY16TMX3Jqo38MG0P8rEAKkLsxnBIA", "random-tenant-id", test.env)
	test.service = services.NewAuthorizationService(test.m)
}

func (test *TestAuthorizationService) TestGetOauthURL() {
	oauthURL := test.service.GetOAuthURL("http://localhost")
	host, _ := url.Parse(oauthURL)

	query := host.Query()

	assert.Equal(test.T(), query.Get("client_id"), test.m.GetClientID())
	assert.Equal(test.T(), query.Get("redirect_uri"), "http://localhost")
}

func (test *TestAuthorizationService) TestIsAuthorizedShouldReturnTrue() {
	tokenService := services.NewAccessTokenService(test.m.GetTenantID(), test.env.GetConfiguration().GetStorage())
	tokenService.Save("key")
	tokenService.SaveExpiration(strconv.Itoa(math.MaxInt))

	isAuthorized := test.service.IsAuthorized()

	assert.True(test.T(), isAuthorized)
}

func (test *TestAuthorizationService) TestIsAuthorizedShouldReturnFalse() {
	tokenService := services.NewAccessTokenService(test.m.GetTenantID(), test.env.GetConfiguration().GetStorage())
	tokenService.Save("key")
	tokenService.SaveExpiration(strconv.Itoa(math.MinInt))

	isAuthorized := test.service.IsAuthorized()

	assert.False(test.T(), isAuthorized)
}

func TestNewAuthorizationService(t *testing.T) {
	suite.Run(t, new(TestAuthorizationService))
}
