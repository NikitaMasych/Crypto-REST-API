package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatCryptoRateAddGetIsSuccessfull(t *testing.T) {
	InitCache()
	expectedRate := 10.213

	AddCurrencyRateToCache(expectedRate)
	receivedRate, err := GetConfigCurrencyRateFromCache()

	assert.Equal(t, err, nil)
	assert.Equal(t, expectedRate, receivedRate)
}
