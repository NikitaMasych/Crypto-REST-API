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
			r.logger.LogDebug("Requesting " + pair.ToString() + " rate from providers chain")
			rate, err = r.providersChain.GetRate(pair)
			if err != nil {
				r.logger.LogError(err)
				return rate, err
			}
			r.cache.AddRateToCache(rate)
			r.logger.LogDebug("Added " + pair.ToString() + " rate to cache")
		}
	} else {
		r.logger.LogDebug("Fetched " + pair.ToString() + " rate from cache")
	}
	return rate, err
}
