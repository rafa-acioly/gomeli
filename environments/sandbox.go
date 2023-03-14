package environments

type sandbox struct {
	wsHost   string
	oauthURI string
	site     site
	config   Configuration
}

// GetSite returns the country that the integration must be done
func (p sandbox) GetSite() site {
	return p.site
}

// GetConfiguration returns the configuration os the resources
func (p sandbox) GetConfiguration() Configuration {
	return p.config
}

// GetWsHost returns the WebService Host endpoint
func (p sandbox) GetWsHost() string {
	return p.wsHost
}

// GetWsAuth returns the WebService Authentication endpoint
func (p sandbox) GetWsAuth() string {
	return wsAuthSandbox[p.GetSite()]
}

// GetWsURL returns the WebService URL with the corresponding URI
// e.g: https://mercadolivre.com.br is the URL and "/items" is the resource
func (p sandbox) GetWsURL(resource string) string {
	return p.GetWsHost() + resource
}

// GetOAuthURI returns the OAuth URI
func (p sandbox) GetOAuthURI() string {
	return p.oauthURI
}

// GetAuthURL returns the authentication endpoint along with the URI
func (p sandbox) GetAuthURL(resource string) string {
	return p.GetWsAuth() + resource
}

func NewSandboxEnv(s site, c Configuration) Environment {
	return &sandbox{
		wsHost:   wsHostSandbox,
		oauthURI: oauthURISandbox,
		site:     s,
		config:   c,
	}
}
