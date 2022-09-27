package services

import "GenesisTask/pkg/domain/models"

type EmailService interface {
	SendRateEmails(models.CurrencyRate, []models.EmailAddress) error
}
