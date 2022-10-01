package crypto

import (
	"GenesisTask/config"
	"GenesisTask/pkg/domain/models"
	"GenesisTask/pkg/platform"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatBinanceProviderReturnsRate(t *testing.T) {
	p := NewBinanceProvider()
	base := "SOL"
	quote := "USDT"
	pair := *models.NewCurrencyPair(base, quote)
	provideArrange()

	rate, err := p.GetRate(pair)

	assert.Equal(t, err, nil)
	assert.NotEqual(t, rate, 0)
}

func TestThatCoinbaseProviderReturnsRate(t *testing.T) {
	p := NewCoinbaseProvider()
	base := "SOL"
	quote := "USDT"
	pair := *models.NewCurrencyPair(base, quote)
	provideArrange()

	rate, err := p.GetRate(pair)

	assert.Equal(t, err, nil)
	assert.NotEqual(t, rate, 0)
}

func TestThatCoinApiProviderReturnsRate(t *testing.T) {
	p := NewCoinApiProvider()
	base := "SOL"
	quote := "USDT"
	pair := *models.NewCurrencyPair(base, quote)
	provideArrange()

	rate, err := p.GetRate(pair)

	assert.Equal(t, err, nil)
	assert.NotEqual(t, rate, 0)
}

func provideArrange() {
	platform.EnsureFileExists(config.LoggerFile)
}
