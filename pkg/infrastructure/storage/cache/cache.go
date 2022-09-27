package cache

import (
	"GenesisTask/config"
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/domain/models"
	"GenesisTask/pkg/errors"
	"time"

	"github.com/patrickmn/go-cache"
)

type GoCache struct {
	cache *cache.Cache
}

func NewGoCache() application.Cache {
	return &GoCache{cache.New(cache.NoExpiration, cache.NoExpiration)}
}

func (g *GoCache) AddRateToCache(rate models.CurrencyRate) {
	existingPeriod := time.Duration(config.Get().CacheDurationMins) * time.Minute
	pair := rate.GetCurrencyPair()
	rateAssets := pair.GetBase() + "-" + pair.GetQuote()
	g.cache.Set(rateAssets, rate.GetPrice(), existingPeriod)
}

func (g *GoCache) GetRateFromCache(pair models.CurrencyPair) (models.CurrencyRate, error) {
	rateAssets := pair.GetBase() + "-" + pair.GetQuote()
	rate, present := g.cache.Get(rateAssets)
	if !present {
		return *models.NewCurrencyRate(pair, -1), errors.ErrNotPresentInCache
	}
	float64Rate, err := convertInterfaceToFloat64(rate)
	return *models.NewCurrencyRate(pair, float64Rate), err
}
