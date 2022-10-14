package crypto

import (
	"producer/config"
	"producer/pkg/application"
	"producer/pkg/domain/models"
	txtLogger "producer/pkg/infrastructure/logger/txt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatBinanceProviderReturnsRate(t *testing.T) {
	p := NewBinanceProvider()
	base := "SOL"
	quote := "USDT"
	pair := *models.NewCurrencyPair(base, quote)
	logger = createTxtLogger()

	rate, err := p.GetRate(pair)

	assert.Equal(t, err, nil)
	assert.NotEqual(t, rate, 0)
}

func TestThatCoinbaseProviderReturnsRate(t *testing.T) {
	p := NewCoinbaseProvider()
	base := "SOL"
	quote := "USDT"
	pair := *models.NewCurrencyPair(base, quote)
	logger = createTxtLogger()

	rate, err := p.GetRate(pair)

	assert.Equal(t, err, nil)
	assert.NotEqual(t, rate, 0)
}

func TestThatCoinApiProviderReturnsRate(t *testing.T) {
	p := NewCoinApiProvider()
	base := "SOL"
	quote := "USDT"
	pair := *models.NewCurrencyPair(base, quote)
	logger = createTxtLogger()

	rate, err := p.GetRate(pair)

	assert.Equal(t, err, nil)
	assert.NotEqual(t, rate, 0)
}

func createTxtLogger() application.Logger {
	loggerFiles := txtLogger.NewLoggerFiles(config.DebugLogFile,
		config.ErrorsLogFile, config.InfoLogFile)
	txtLogger.EnsureLogFilesExist(loggerFiles)
	logger := txtLogger.NewTxtLogger(loggerFiles)
	return logger
}
