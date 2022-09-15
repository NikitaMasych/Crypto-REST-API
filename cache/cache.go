package cache

import (
	"GenesisTask/config"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/allegro/bigcache"
)

var (
	cache *bigcache.BigCache
	once  sync.Once
)

func InitCache() {
	once.Do(func() {
		var err error
		duration := time.Duration(config.Get().CacheDurationMins)
		cache, err = bigcache.NewBigCache(bigcache.DefaultConfig(duration))
		if err != nil {
			log.Fatal(err)
		}
	})
}

func AddCryptoRateToCache(price float64) {
	cfg := config.Get()
	rateAssets := cfg.BaseCurrency + "-" + cfg.QuotedCurrency
	convertedPrice := []byte(fmt.Sprintf("%f", price))
	err := cache.Set(rateAssets, convertedPrice)
	if err != nil {
		log.Fatal(err)
	}
}

func GetCryptoRateFromCache() (float64, error) {
	cfg := config.Get()
	rateAssets := cfg.BaseCurrency + "-" + cfg.QuotedCurrency
	bytePrice, err := cache.Get(rateAssets)
	if err != nil {
		return 0, err
	}
	stringPrice := string(bytePrice)
	price, err := strconv.ParseFloat(stringPrice, 64)
	if err != nil {
		return 0, err
	}
	return price, nil
}
