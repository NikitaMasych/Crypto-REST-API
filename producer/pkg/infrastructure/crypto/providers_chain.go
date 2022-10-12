package crypto

import (
	"producer/config"
	"producer/pkg/application"
)

func NewProvidersChain(l application.Logger) *application.ProvidersChain {
	logger = l
	coinbaseProvider := NewCoinbaseProvider()
	l.LogDebug("Created coinbase provider")
	binanceProvider := NewBinanceProvider()
	l.LogDebug("Created binance provider")
	coinApiProvider := NewCoinApiProvider()
	l.LogDebug("Created coinApi provider")

	var genesisProvider application.ProvidersChain
	switch config.GenesisProvider {
	case "binance":
		genesisProvider = binanceProvider
		binanceProvider.SetNext(coinbaseProvider)
		coinbaseProvider.SetNext(coinApiProvider)
		l.LogDebug("Initialized chain in order: binance, coinbase, coinapi")
	case "coinbase":
		genesisProvider = coinbaseProvider
		coinbaseProvider.SetNext(binanceProvider)
		binanceProvider.SetNext(coinApiProvider)
		l.LogDebug("Initialized chain in order: coinbase, binance, coinapi")
	case "coinapi":
		genesisProvider = coinApiProvider
		coinApiProvider.SetNext(binanceProvider)
		binanceProvider.SetNext(coinbaseProvider)
		l.LogDebug("Initialized chain in order: coinapi, binance, coinbase")
	default:
		genesisProvider = binanceProvider
		binanceProvider.SetNext(coinbaseProvider)
		coinbaseProvider.SetNext(coinApiProvider)
		l.LogDebug("Initialized chain in order: binance, coinbase, coinapi")
	}
	return &genesisProvider
}
