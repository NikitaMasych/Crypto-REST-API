package application

import (
	"GenesisTask/pkg/domain/models"
	custom "GenesisTask/pkg/errors"
	"errors"
)

type RateRepository struct {
	providersChain ProvidersChain
	cache          Cache
	logger         Logger
}

func NewRateRepository(providersChain ProvidersChain,
	cache Cache, logger Logger) *RateRepository {
	return &RateRepository{providersChain, cache, logger}
}

func (r *RateRepository) GetRate(pair models.CurrencyPair) (models.CurrencyRate, error) {
	rate, err := r.cache.GetRateFromCache(pair)
	if err != nil {
		r.logger.LogError(err)
		if errors.Is(err, custom.ErrNotPresentInCache) {
			rate, err = r.providersChain.GetRate(pair)
			if err != nil {
				r.logger.LogError(err)
				return rate, err
			}
			r.cache.AddRateToCache(rate)
		}
	}
	return rate, err
}
