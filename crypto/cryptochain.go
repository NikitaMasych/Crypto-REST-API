package crypto

import (
	"GenesisTask/cache"
	"GenesisTask/config"
)

type Cryptochain interface {
	SetNext(next *Cryptochain)
	GetCurrencyRate(base, quoted string) (float64, error)
}

func NewCryptochain() Cryptochain {
	service1 := NewCoinbaseService()
	service2 := NewBinanceService()
	service3 := NewCoinApiService()

	service1.SetNext(&service2)
	service2.SetNext(&service3)

	return service1
}

func GetConfigCurrencyRate() (float64, error) {
	cfg := config.Get()
	c := NewCryptochain()
	rate, err := c.GetCurrencyRate(cfg.BaseCurrency, cfg.QuotedCurrency)
	if err != nil {
		return 0, err
	}

	cache.AddCurrencyRateToCache(rate)

	return rate, err
}
