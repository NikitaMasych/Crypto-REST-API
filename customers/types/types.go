package types

import "gorm.io/gorm"

type (
	Customer struct {
		gorm.Model
		EmailAddress string
	}
	Order struct {
		gorm.Model
		IDTransaction string
		EmailAddress  string
		Status        string
	}
)
