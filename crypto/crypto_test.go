package crypto

import (
	"GenesisTask/config"
	"GenesisTask/platform"
	"os"
	"strings"
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

	cleanup(t)
}

func TestThatCoinbaseProviderReturnsRate(t *testing.T) {
	var p CoinbaseProvider
	cfg := config.Get()
	provideArrange()

	rate, err := p.getCurrencyRate(cfg.BaseCurrency, cfg.QuotedCurrency)

	assert.Equal(t, err, nil)
	assert.NotEqual(t, rate, 0)

	cleanup(t)
}

func TestThatCoinApiProviderReturnsRate(t *testing.T) {
	var p CoinApiProvider
	cfg := config.Get()
	provideArrange()

	rate, err := p.getCurrencyRate(cfg.BaseCurrency, cfg.QuotedCurrency)

	assert.Equal(t, err, nil)
	assert.NotEqual(t, rate, 0)

	cleanup(t)
}

func provideArrange() {
	platform.EnsureFileExists(config.Get().LoggerFile)
}

func cleanup(t *testing.T) {
	path := config.Get().LoggerFile
	_, err := os.Stat(path)
	if err != nil {
		t.Error(err)
	}
	directory := strings.Split(path, "/")[0]
	err = os.RemoveAll(directory)
	if err != nil {
		t.Error(err)
	}
}
