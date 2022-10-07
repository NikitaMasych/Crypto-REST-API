package application

import (
	"GenesisTask/pkg/domain/models"
	"GenesisTask/pkg/domain/services"
)

type ProvidersChain interface {
	services.RateService
	SetNext(ProvidersChain)
}

type Cache interface {
	AddRateToCache(models.CurrencyRate)
	GetRateFromCache(models.CurrencyPair) (models.CurrencyRate, error)
}

type SubscriptionStorage interface {
	AddSubscription(models.Subscription) error
	IsSaved(models.Subscription) bool
	GetAll() []models.Subscription
}

type EmailSender interface {
	SendRatesEmail([]models.CurrencyRate, models.EmailAddress)
}

type Logger interface {
	loggerDebug
	loggerError
	loggerInfo
}

type loggerDebug interface {
	LogDebug(v ...any)
}

type loggerError interface {
	LogError(v ...any)
}

type loggerInfo interface {
	LogInfo(v ...any)
}
