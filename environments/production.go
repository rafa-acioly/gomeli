package environments

type Environment interface {
	GetWsHost() string
	GetOAuthURI() string
	EnvironmentConfig
}

type production struct {
	wsHost   string
	oauthURI string
	env      EnvironmentConfig
}

func (p production) GetSite() site {
	return p.env.GetSite()
}

func (p production) GetConfiguration() Configuration {
	return p.env.GetConfiguration()
}

func (p production) GetWsHost() string {
	return p.wsHost
}

func (p production) GetOAuthURI() string {
	return p.oauthURI
}

func NewProductionEnv(s site, c Configuration) Environment {
	return &production{
		wsHost:   wsHostProduction,
		oauthURI: oauthURIProduction,
		env:      NewEnvironmentConfig(s, c),
	}
}
