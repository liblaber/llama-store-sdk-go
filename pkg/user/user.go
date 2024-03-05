package user

// A user of the llama store
type User struct {
	// The email address of the user. This must be unique across all users.
	Email *string `json:"email,omitempty" required:"true"`
	// The ID of the user. This is unique across all users.
	Id *int64 `json:"id,omitempty" required:"true"`
}

func (u *User) SetEmail(email string) {
	u.Email = &email
}

func (u *User) GetEmail() *string {
	if u == nil {
		return nil
	}
	return u.Email
}

func (u *User) SetId(id int64) {
	u.Id = &id
}

func (u *User) GetId() *int64 {
	if u == nil {
		return nil
	}
	return u.Id
}
