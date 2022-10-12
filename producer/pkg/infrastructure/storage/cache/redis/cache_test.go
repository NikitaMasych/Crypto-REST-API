package redis

import (
	"os"
	"producer/config"
	"producer/pkg/domain/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestThatCryptoRateAddGetIsSuccessfull(t *testing.T) {
	os.Setenv("REDIS_CACHE_HOST", "redis-cache:6379")
	c := NewRedisCache(config.CacheDb,
		time.Duration(config.CacheDurationMins)*time.Minute)
	base := "SOL"
	quote := "USDT"
	pair := *models.NewCurrencyPair(base, quote)
	expectedRate := 35.323
	timestamp := time.Now()
	rate := models.NewCurrencyRate(pair, expectedRate, timestamp)

	c.AddRateToCache(*rate)
	receivedRate, err := c.GetRateFromCache(pair)

	assert.Equal(t, err, nil)
	assert.Equal(t, expectedRate, receivedRate.GetPrice())
	assert.Equal(t, timestamp, receivedRate.GetTimestamp())
}
