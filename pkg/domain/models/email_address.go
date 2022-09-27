package models

type EmailAddress struct {
	address string
}

func NewEmailAddress(address string) *EmailAddress {
	return &EmailAddress{address}
}

func (e *EmailAddress) GetAddress() string {
	return e.address
}
