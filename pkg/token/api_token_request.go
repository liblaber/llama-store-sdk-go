package token

// A request to get an API token for a given user.
type ApiTokenRequest struct {
	// The email address of the user. This must be unique across all users.
	Email *string `json:"email,omitempty" required:"true"`
	// The password of the user. This must be at least 8 characters long, and contain at least one letter, one number, and one special character.
	Password *string `json:"password,omitempty" required:"true"`
}

func (a *ApiTokenRequest) SetEmail(email string) {
	a.Email = &email
}

func (a *ApiTokenRequest) GetEmail() *string {
	if a == nil {
		return nil
	}
	return a.Email
}

func (a *ApiTokenRequest) SetPassword(password string) {
	a.Password = &password
}

func (a *ApiTokenRequest) GetPassword() *string {
	if a == nil {
		return nil
	}
	return a.Password
}
