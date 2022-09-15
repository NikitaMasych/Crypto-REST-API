package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsConfigPresent(t *testing.T) {
	cfg := Get()

	require.NotEmpty(t, cfg.ServerURL)
	require.NotEmpty(t, cfg.BinanceApiFormatUrl)
	require.NotEmpty(t, cfg.CoinbaseApiFormatUrl)
	require.NotEmpty(t, cfg.CoinApiFormatURL)
	require.NotEmpty(t, cfg.BaseCurrency)
	require.NotEmpty(t, cfg.QuotedCurrency)
	require.NotEmpty(t, cfg.CryptoCurrencyProvider)
	require.NotEmpty(t, cfg.EmailAddress)
	require.NotEmpty(t, cfg.EmailPassword)
	require.NotEmpty(t, cfg.StorageFile)
	require.NotEmpty(t, cfg.LoggerFile)
	require.NotEmpty(t, cfg.SMTPHost)
	assert.NotEqual(t, cfg.SMTPPort, 0)
	assert.NotEqual(t, cfg.CacheDurationMins, 0)
}
