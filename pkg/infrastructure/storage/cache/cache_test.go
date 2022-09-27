package cache

import (
	"GenesisTask/config"
	"GenesisTask/pkg/domain/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatCryptoRateAddGetIsSuccessfull(t *testing.T) {
	c := NewGoCache()
	s := config.NewConfigPairSource()
	pair := s.GetPair()
	expectedRate := 10.213
	rate := models.NewCurrencyRate(pair, expectedRate)

	c.AddRateToCache(*rate)
	receivedRate, err := c.GetRateFromCache(pair)

	assert.Equal(t, err, nil)
	assert.Equal(t, expectedRate, receivedRate.GetPrice())
}
