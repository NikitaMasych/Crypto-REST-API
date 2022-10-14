package types

import "gorm.io/gorm"

type (
	Customer struct {
		gorm.Model
		CustomerId   string
		EmailAddress string
	}
	Order struct {
		gorm.Model
		TransactionId string
		CustomerId    string
		EmailAddress  string
		Status        string
	}
)
