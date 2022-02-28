package environments

type sandbox struct {
	wsHost   string
	oauthURI string
	env      EnvironmentConfig
}

func (s sandbox) GetWsHost() string {
	return s.wsHost
}

func (s sandbox) GetOAuthURI() string {
	return s.oauthURI
}

func (s sandbox) GetSite() site {
	return s.env.GetSite()
}

func (s sandbox) GetConfiguration() Configuration {
	return s.env.GetConfiguration()
}

func NewSandboxEnv(s site, c Configuration) Environment {
	return &sandbox{
		wsHost:   wsHostSandbox,
		oauthURI: oauthURISandbox,
		env:      NewEnvironmentConfig(s, c),
	}
}
