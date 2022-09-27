package application

import (
	"GenesisTask/pkg/domain/models"
	custom "GenesisTask/pkg/errors"
	"errors"
)

type RateRepository struct {
	providersChain ProvidersChain
	cache          Cache
}

func NewRateRepository(providersChain ProvidersChain,
	cache Cache) *RateRepository {
	return &RateRepository{providersChain, cache}
}

func (r *RateRepository) GetRate(pair models.CurrencyPair) (models.CurrencyRate, error) {
	rate, err := r.cache.GetRateFromCache(pair)
	if err != nil {
		if errors.Is(err, custom.ErrNotPresentInCache) {
			rate, err = r.providersChain.GetRate(pair)
			if err != nil {
				return rate, err
			}
			r.cache.AddRateToCache(rate)
		}
	}
	return rate, err
}
