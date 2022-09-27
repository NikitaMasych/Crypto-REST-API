package models

type User struct {
	EmailAddress
}

func NewUser(e EmailAddress) *User {
	return &User{e}
}

func (u *User) GetEmailAddress() *EmailAddress {
	return &u.EmailAddress
}
