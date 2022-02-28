package meli

type Meli interface {
	GetClientID() string
	GetClientSecret() string
	GetUserID() string
	GetTenantID() string
	GetEnvironment() string
}

type meli struct {
	clientID     string
	clientSecret string
	userID       string
	tenantID     string
	environment  string
}

func (m meli) GetClientID() string {
	return m.clientID
}

func (m meli) GetClientSecret() string {
	return m.clientSecret
}

func (m meli) GetUserID() string {
	return m.userID
}

func (m meli) GetTenantID() string {
	return m.tenantID
}

func (m meli) GetEnvironment() string {
	return m.environment
}

func NewMeli(clientID, clientSecret string) Meli {
	return &meli{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}
