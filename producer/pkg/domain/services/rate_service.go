package services

import "producer/pkg/domain/models"

type RateService interface {
	GetRate(pair models.CurrencyPair) (models.CurrencyRate, error)
}
