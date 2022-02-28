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

func (s sandbox) GetWsAuth() map[site]string {
	return map[site]string{
		ARGENTINA:  "https://auth.mercadolibre.com.ar",
		BOLIVIA:    "https://auth.mercadolibre.com.bo",
		BRASIL:     "https://auth.mercadolivre.com.br",
		COLOMBIA:   "https://auth.mercadolibre.com.co",
		COSTA_RICA: "https://auth.mercadolibre.com.cr",
		EQUADOR:    "https://auth.mercadolibre.com.ec",
		HONDURAS:   "https://auth.mercadolibre.com.hn",
		GUATEMALA:  "https://auth.mercadolibre.com.gt",
		CHILE:      "https://auth.mercadolibre.cl",
		MEXICO:     "https://auth.mercadolibre.com.mx",
		NICARAGUA:  "https://auth.mercadolibre.com.ni",
		PARAGUAI:   "https://auth.mercadolibre.com.py",
		SALVADOR:   "https://auth.mercadolibre.com.sv",
		URUGUAI:    "https://auth.mercadolibre.com.uy",
		VENEZUELA:  "https://auth.mercadolibre.com.ve",
		PANAMA:     "https://auth.mercadolibre.com.pa",
		PERU:       "https://auth.mercadolibre.com.pe",
		PORTUGAL:   "https://auth.mercadolibre.com.pt",
		DOMINICANA: "https://auth.mercadolibre.com.do",
	}
}

func (s sandbox) GetSite() site {
	return s.env.GetSite()
}

func (s sandbox) GetConfiguration() Configuration {
	return s.env.GetConfiguration()
}

func NewSandboxEnv(s site, c Configuration) Environment {
	return &sandbox{
		wsHost:   "https://api.mercadolibre.com",
		oauthURI: "/oauth/token",
		env:      NewEnvironmentConfig(s, c),
	}
}
