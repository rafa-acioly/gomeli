package environments

// Environment defines the contract that need to be used for each environment
// e.g: production, sandbox, staging and etc
type Environment interface {
	GetWsHost() string
	GetWsAuth() string
	GetWsURL(resource string) string
	GetOAuthURI() string
	GetAuthURL(resource string) string
	GetSite() site
	GetConfiguration() Configuration
}

type production struct {
	wsHost   string
	oauthURI string
	site     site
	config   Configuration
}

// GetSite returns the country that the integration must be done
func (p production) GetSite() site {
	return p.GetSite()
}

// GetConfiguration returns the configuration os the resources
func (p production) GetConfiguration() Configuration {
	return p.GetConfiguration()
}

// GetWsHost returns the WebService Host endpoint
func (p production) GetWsHost() string {
	return p.wsHost
}

// GetWsAuth returns the WebService Authentication endpoint
func (p production) GetWsAuth() string {
	return wsAuthProduction[p.GetSite()]
}

// GetWsURL returns the WebService URL with the corresponding URI
// e.g: https://mercadolivre.com.br is the URL and "/items" is the resource
func (p production) GetWsURL(resource string) string {
	return p.GetWsHost() + resource
}

// GetOAuthURI returns the OAuth URI
func (p production) GetOAuthURI() string {
	return p.oauthURI
}

// GetAuthURL returns the authentication endpoint along with the URI
func (p production) GetAuthURL(resource string) string {
	return p.GetWsAuth() + resource
}

// NewProductionEnv returns the configuration to use on production mode
func NewProductionEnv(s site, c Configuration) Environment {
	return &production{
		wsHost:   wsHostProduction,
		oauthURI: oauthURIProduction,
		site:     s,
		config:   c,
	}
}
