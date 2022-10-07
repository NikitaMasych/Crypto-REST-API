package models

type EmailAddress struct {
	address string
}

func NewEmailAddress(address string) *EmailAddress {
	return &EmailAddress{address}
}

func (e *EmailAddress) ToString() string {
	return e.address
}
