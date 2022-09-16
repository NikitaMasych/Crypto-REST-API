package cache

import (
	"GenesisTask/config"
	"GenesisTask/errors"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	c    *cache.Cache
	once sync.Once
)

func InitCache() {
	once.Do(func() {
		c = cache.New(cache.NoExpiration, cache.NoExpiration)
	})
}

func AddCurrencyRateToCache(price float64) {
	cfg := config.Get()
	rateAssets := cfg.BaseCurrency + "-" + cfg.QuotedCurrency
	existingPeriod := time.Duration(config.Get().CacheDurationMins) * time.Minute
	c.Set(rateAssets, price, existingPeriod)
}

func GetConfigCurrencyRateFromCache() (float64, error) {
	cfg := config.Get()
	rateAssets := cfg.BaseCurrency + "-" + cfg.QuotedCurrency
	rate, present := c.Get(rateAssets)
	if !present {
		return 0, errors.ErrNotPresentInCache
	}

	return convertInterfaceToFloat64(rate)
}
