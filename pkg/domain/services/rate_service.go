package services

import "GenesisTask/pkg/domain/models"

type RateService interface {
	GetRate(pair models.CurrencyPair) (models.CurrencyRate, error)
}
