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
