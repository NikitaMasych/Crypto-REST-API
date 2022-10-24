package customers

import (
	"producer/pkg/application"
	"producer/pkg/domain/models"

	"gopkg.in/resty.v0"
)

type CustomersService struct {
	customerCreationURL string
}

func NewCustomersService(customerCreationURL string) application.CustomersService {
	return &CustomersService{customerCreationURL}
}

func (c *CustomersService) CreateCustomer(emailAddress models.EmailAddress) error {
	body := struct {
		EmailAddress string `json:"emailAddress"`
	}{emailAddress.ToString()}

	_, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(c.customerCreationURL)

	return err
}
