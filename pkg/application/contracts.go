package application

import (
	"GenesisTask/pkg/domain/models"
	"GenesisTask/pkg/domain/services"
)

type ProvidersChain interface {
	services.RateService
	SetNext(*ProvidersChain)
}

type Cache interface {
	AddRateToCache(models.CurrencyRate)
	GetRateFromCache(models.CurrencyPair) (models.CurrencyRate, error)
}

type EmailAddressesStorage interface {
	AddEmail(models.EmailAddress) error
	IsSaved(models.EmailAddress) bool
	GetAll() []models.EmailAddress
}

type EmailSender interface {
	SendRateEmails(models.CurrencyRate, []models.EmailAddress)
}

type PairSource interface {
	GetPair() models.CurrencyPair
}
