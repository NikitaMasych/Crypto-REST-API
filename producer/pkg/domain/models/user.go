package models

type User struct {
	EmailAddress
	SubscribedPairs []CurrencyPair
}

func NewUser(e EmailAddress, p []CurrencyPair) *User {
	return &User{e, p}
}

func (u *User) GetEmailAddress() *EmailAddress {
	return &u.EmailAddress
}

func (u *User) GetSubscribedPairs() []CurrencyPair {
	return u.SubscribedPairs
}
