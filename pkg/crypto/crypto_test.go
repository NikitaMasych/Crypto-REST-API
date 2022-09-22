package crypto

import (
	"GenesisTask/config"
	"GenesisTask/pkg/platform"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatBinanceProviderReturnsRate(t *testing.T) {
	var p BinanceProvider
	cfg := config.Get()
	provideArrange()

	rate, err := p.getCurrencyRate(cfg.BaseCurrency, cfg.QuotedCurrency)

	assert.Equal(t, err, nil)
	assert.NotEqual(t, rate, 0)
}

func TestThatCoinbaseProviderReturnsRate(t *testing.T) {
	var p CoinbaseProvider
	cfg := config.Get()
	provideArrange()

	rate, err := p.getCurrencyRate(cfg.BaseCurrency, cfg.QuotedCurrency)

	assert.Equal(t, err, nil)
	assert.NotEqual(t, rate, 0)
}

func TestThatCoinApiProviderReturnsRate(t *testing.T) {
	var p CoinApiProvider
	cfg := config.Get()
	provideArrange()

	rate, err := p.getCurrencyRate(cfg.BaseCurrency, cfg.QuotedCurrency)

	assert.Equal(t, err, nil)
	assert.NotEqual(t, rate, 0)
}

func provideArrange() {
	platform.EnsureFileExists(config.Get().LoggerFile)
}
