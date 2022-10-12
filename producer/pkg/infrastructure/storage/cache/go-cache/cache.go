package gocache

import (
	"producer/pkg/application"
	"producer/pkg/domain/models"
	"producer/pkg/errors"
	"time"

	"github.com/patrickmn/go-cache"
)

type GoCache struct {
	cache           *cache.Cache
	ratesExpiration time.Duration
}

func NewGoCache(ratesExpiration time.Duration) application.Cache {
	return &GoCache{cache.New(cache.NoExpiration, cache.NoExpiration), ratesExpiration}
}

type rateTrait struct {
	Rate      float64   `json:"rate"`
	Timestamp time.Time `json:"timestamp"`
}

func (g *GoCache) AddRateToCache(rate models.CurrencyRate) {
	pair := rate.GetCurrencyPair()
	rateAssets := pair.GetBase() + "-" + pair.GetQuote()
	rateStruct := rateTrait{rate.GetPrice(), rate.GetTimestamp()}
	g.cache.Set(rateAssets, rateStruct, time.Duration(g.ratesExpiration))
}

func (g *GoCache) GetRateFromCache(pair models.CurrencyPair) (models.CurrencyRate, error) {
	rateAssets := pair.GetBase() + "-" + pair.GetQuote()
	rate, present := g.cache.Get(rateAssets)
	if !present {
		return *models.NewCurrencyRate(pair, -1, time.Now()), errors.ErrNotPresentInCache
	}
	result := rate.(rateTrait)
	return *models.NewCurrencyRate(pair, result.Rate, result.Timestamp), nil
}
