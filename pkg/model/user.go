package model

type User struct {
	email string
}

func NewUser(email string) *User {
	return &User{email: email}
}

func (u *User) GetEmail() string {
	return u.email
}
