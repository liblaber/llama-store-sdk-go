package user

// A new user of the llama store.
type UserRegistration struct {
	// The email address of the user. This must be unique across all users.
	Email *string `json:"email,omitempty" required:"true"`
	// The password of the user. This must be at least 8 characters long, and contain at least one letter, one number, and one special character.
	Password *string `json:"password,omitempty" required:"true"`
}

func (u *UserRegistration) SetEmail(email string) {
	u.Email = &email
}

func (u *UserRegistration) GetEmail() *string {
	if u == nil {
		return nil
	}
	return u.Email
}

func (u *UserRegistration) SetPassword(password string) {
	u.Password = &password
}

func (u *UserRegistration) GetPassword() *string {
	if u == nil {
		return nil
	}
	return u.Password
}
