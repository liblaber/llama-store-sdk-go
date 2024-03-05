package token

// An API token to use for authentication.
type ApiToken struct {
	// The bearer token to use with the API. Pass this in the Authorization header as a bearer token.
	AccessToken *string `json:"access_token,omitempty" required:"true"`
	// The type of token. This will always be bearer.
	TokenType *string `json:"token_type,omitempty"`
}

func (a *ApiToken) SetAccessToken(accessToken string) {
	a.AccessToken = &accessToken
}

func (a *ApiToken) GetAccessToken() *string {
	if a == nil {
		return nil
	}
	return a.AccessToken
}

func (a *ApiToken) SetTokenType(tokenType string) {
	a.TokenType = &tokenType
}

func (a *ApiToken) GetTokenType() *string {
	if a == nil {
		return nil
	}
	return a.TokenType
}
