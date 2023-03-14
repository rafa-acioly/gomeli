package services

import (
	"fmt"
	"meli"
	"meli/responses"
	"net/http"
	"net/url"
	"strconv"
)

// AuthorizationService represents the contract to access
// authorization API on mercado livre's API
type AuthorizationService interface {
	GetOAuthURL(redirectURI string) string
	GetAuthorizationCode(redirectURI string) string
	Authorize(code, redirectURI string) (string, error)
	GetAccessToken() (string, error)
	IsAuthorized() bool
	getToken(url string, values url.Values) (responses.Authorization, responses.AuthorizationError, error)
}

type authorizationService struct {
	http meli.HttpClientWrapper
	meli meli.Meli
}

// Authorize exchange the code given by mercado livre to an useful access token,
// this code can be retrieved when the user allow the Oauth2 integration
// using the method GetOAuthURL
func (auth authorizationService) Authorize(code string, redirectURI string) (string, error) {
	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("client_id", auth.meli.GetClientID())
	values.Add("client_secret", auth.meli.GetClientSecret())
	values.Add("code", code)
	values.Add("redirect_uri", redirectURI)

	environment := auth.meli.GetEnvironment()
	success, fail, err := auth.getToken(environment.GetOAuthURI(), values)
	if err != nil {
		return "", fmt.Errorf("could not retrieve token error=%s message=%s", err.Error(), fail.Message)
	}

	tokenService := NewAccessTokenService(
		auth.meli.GetTenantID(),
		auth.meli.GetEnvironment().GetConfiguration().GetStorage(),
	)
	tokenService.Save(success.AccessToken)
	tokenService.SaveExpiration(strconv.Itoa(success.ExpireIn))

	return success.AccessToken, nil
}

// GetAccessToken return the access token, this form is recommended for scripts that run in automatic
// routines (via cron, or scheduled tasks). OBS: to be able to use it, you need to have
// Scope offline access checked in your APP
func (auth authorizationService) GetAccessToken() (string, error) {
	values := url.Values{}
	values.Add("grant_type", "client_credentials")
	values.Add("client_id", auth.meli.GetClientID())
	values.Add("client_secret", auth.meli.GetClientSecret())

	environment := auth.meli.GetEnvironment()
	success, fail, err := auth.getToken(environment.GetOAuthURI(), values)
	if err != nil {
		return "", fmt.Errorf("could not retrieve token error=%s message=%s", err.Error(), fail.Message)
	}

	tokenService := NewAccessTokenService(
		auth.meli.GetTenantID(),
		auth.meli.GetEnvironment().GetConfiguration().GetStorage(),
	)
	tokenService.Save(success.AccessToken)
	tokenService.SaveExpiration(strconv.Itoa(success.ExpireIn))

	return success.AccessToken, nil
}

// GetAuthorizationCode implements AuthorizationService
func (auth authorizationService) GetAuthorizationCode(redirectURI string) (string, error) {
	values := url.Values{}
	values.Add("grant_type", "code")
	values.Add("client_id", auth.meli.GetClientID())
	values.Add("redirect_uri", redirectURI)

	environment := auth.meli.GetEnvironment()
	success, fail, err := auth.getToken(environment.GetWsAuth(), values)
	if err != nil {
		return "", fmt.Errorf("could not retrieve token error=%s message=%s", err.Error(), fail.Message)
	}

	return success.AccessToken, nil
}

// GetOAuthURL return the URL used to authorize Oauth2 integration,
// this URL will lead to mercado livre's page asking for
// an user authorization to use your app
func (auth *authorizationService) GetOAuthURL(redirectURI string) string {
	environment := auth.meli.GetEnvironment()

	host, _ := url.Parse(environment.GetAuthURL("/authorization"))

	parameters := host.Query()
	parameters.Add("client_id", auth.meli.GetClientID())
	parameters.Add("response_type", "code")
	parameters.Add("redirect_uri", redirectURI)

	host.RawQuery = parameters.Encode()

	return host.String()
}

// IsAuthorized check is further requests will be authorized
// using the current authorization configuration
func (auth *authorizationService) IsAuthorized() bool {
	environment := auth.meli.GetEnvironment()
	storage := environment.GetConfiguration().GetStorage()
	tenantID := auth.meli.GetTenantID()

	accessTokenService := NewAccessTokenService(tenantID, storage)

	return accessTokenService.IsValid()
}

func (auth *authorizationService) getToken(url string, values url.Values) (responses.Authorization, responses.AuthorizationError, error) {
	authorization := responses.Authorization{}
	authorizationError := responses.AuthorizationError{}
	response, err := auth.http.
		SetHeader("skipOAuth", "true").
		SetResult(&authorization).
		SetError(&authorizationError).
		Post(url, values.Encode())

	if err != nil || response.StatusCode() != http.StatusOK {
		errorContent := fmt.Errorf(
			"could not retrieve token, status=%d message=%s",
			response.StatusCode(),
			authorizationError.Message,
		)
		return authorization, authorizationError, errorContent
	}

	return authorization, authorizationError, nil
}

// NewAuthorizationService is responsible for handling authorization requests,
// and retrieving access tokens
func NewAuthorizationService(m meli.Meli) AuthorizationService {
	return &authorizationService{meli: m, http: meli.NewHttpClient(m)}
}
