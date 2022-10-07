package gocache

import (
	"GenesisTask/config"
	"GenesisTask/pkg/domain/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestThatCryptoRateAddGetIsSuccessfull(t *testing.T) {
	cacheDuration := time.Duration(config.CacheDurationMins) * time.Minute
	c := NewGoCache(cacheDuration)
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
