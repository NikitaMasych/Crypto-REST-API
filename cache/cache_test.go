package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatCryptoRateAddGetIsSuccessfull(t *testing.T) {
	InitCache()
	expectedRate := 10.213

	AddCryptoRateToCache(expectedRate)
	receivedRate, err := GetCryptoRateFromCache()

	assert.Equal(t, err, nil)
	assert.Equal(t, expectedRate, receivedRate)
}
