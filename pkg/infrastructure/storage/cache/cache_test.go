package cache

import (
	"GenesisTask/config"
	"GenesisTask/pkg/domain/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestThatCryptoRateAddGetIsSuccessfull(t *testing.T) {
	c := NewGoCache()
	s := config.NewConfigPairSource()
	pair := s.GetPair()
	expectedRate := 10.213
	timestamp := time.Now()
	rate := models.NewCurrencyRate(pair, expectedRate, timestamp)

	c.AddRateToCache(*rate)
	receivedRate, err := c.GetRateFromCache(pair)

	assert.Equal(t, err, nil)
	assert.Equal(t, expectedRate, receivedRate.GetPrice())
	assert.Equal(t, timestamp, receivedRate.GetTimestamp())
}
