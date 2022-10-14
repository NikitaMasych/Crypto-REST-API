package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"os"
	"producer/pkg/application"
	"producer/pkg/domain/models"

	"gopkg.in/resty.v0"
)

var customerCreationURL = os.Getenv("CUSTOMER_CREATION_URL")

func CreateCustomer(emailAddress models.EmailAddress, logger application.Logger) {
	hasher := sha256.New()
	hasher.Write([]byte(emailAddress.ToString()))
	customerId := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	body := struct {
		CustomerId   string `json:customerId`
		EmailAddress string `json:"emailAddress"`
	}{customerId, emailAddress.ToString()}

	_, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(customerCreationURL)

	if err != nil {
		logger.LogError(err)
	}
}
