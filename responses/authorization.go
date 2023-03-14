package responses

// Authorization represents the content retrieved by
// mercado livre's API when asking for
// authorization data
type Authorization struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpireIn    int    `json:"expire_in"`
	Scope       string `json:"scope"`
	UserID      int    `json:"user_id"`
}

type AuthorizationError struct {
	Message string        `json:"message"`
	Error   string        `json:"error"`
	Status  int           `json:"status"`
	Cause   []interface{} `json:"cause"`
}
