package services

import (
	"encoding/json"
	"fmt"
	"io"
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
	getToken(map[string]string) (string, error)
}

type authorizationService struct {
	http meli.HttpClientWrapper
	meli meli.Meli
}

// Authorize exchange the code by a access token.
func (auth authorizationService) Authorize(code string, redirectURI string) (string, error) {
	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("client_id", auth.meli.GetClientID())
	values.Add("client_secret", auth.meli.GetClientSecret())
	values.Add("code", code)
	values.Add("redirect_uri", redirectURI)

	environment := auth.meli.GetEnvironment()
	authorization := responses.Authorization{}
	authorizationError := responses.AuthorizationError{}
	response, err := auth.http.
		SetHeader("skipOAuth", "true").
		SetResult(&authorization).
		SetError(&authorizationError).
		Post(environment.GetOAuthURI(), values.Encode())

	if err != nil || response.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("could not retrieve token error=%s message=%s", err.Error(), authorizationError.Message)
	}

	tokenService := NewAccessTokenService(
		auth.meli.GetTenantID(),
		auth.meli.GetEnvironment().GetConfiguration().GetStorage(),
	)
	tokenService.Save(authorization.AccessToken)
	tokenService.SaveExpiration(strconv.Itoa(authorization.ExpireIn))

	return authorization.AccessToken, nil
}

// GetAccessToken return the access token, this form is recommended for scripts that run in automatic
// routines (via cron, or scheduled tasks). OBS: to be able to use it, you need to have
// Scope offline access checked in your APP.
func (auth authorizationService) GetAccessToken() (string, error) {
	parameters := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     auth.meli.GetClientID(),
		"client_secret": auth.meli.GetClientSecret(),
	}

	return auth.getToken(parameters)
}

// GetAuthorizationCode implements AuthorizationService
func (auth authorizationService) GetAuthorizationCode(redirectURI string) string {
	panic("unimplemented")
}

// GetOAuthURL return the URL used to authorize Oauth2 integration,
// this URL will lead to mercado livre's page asking for
// an user authorization to use your app.
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

func (auth *authorizationService) getToken(data map[string]string) (string, error) {
	environment := auth.meli.GetEnvironment()
	uri := environment.GetOAuthURI()

	response, err := http.Post(uri, "application/json", nil)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	authorization := &responses.Authorization{}
	if err := json.Unmarshal(body, authorization); err != nil {
		return "", err
	}

	storage := environment.GetConfiguration().GetStorage()
	accessTokenService := NewAccessTokenService(auth.meli.GetTenantID(), storage)
	accessTokenService.Save(authorization.AccessToken)
	accessTokenService.SaveRefreshToken(authorization.RefreshToken)
	accessTokenService.SaveExpiration(strconv.Itoa(authorization.ExpireIn))

	return authorization.AccessToken, nil
}

func NewAuthorizationService(m meli.Meli) AuthorizationService {
	return &authorizationService{meli: m, http: meli.NewHttpClient(m)}
}
