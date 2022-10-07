package crypto

import (
	"GenesisTask/config"
	"GenesisTask/pkg/application"
)

func NewProvidersChain(l application.Logger) *application.ProvidersChain {
	logger = l
	coinbaseProvider := NewCoinbaseProvider()
	binanceProvider := NewBinanceProvider()
	coinApiProvider := NewCoinApiProvider()

	var genesisProvider application.ProvidersChain
	switch config.GenesisProvider {
	case "binance":
		genesisProvider = binanceProvider
		binanceProvider.SetNext(coinbaseProvider)
		coinbaseProvider.SetNext(coinApiProvider)
	case "coinbase":
		genesisProvider = coinbaseProvider
		coinbaseProvider.SetNext(binanceProvider)
		binanceProvider.SetNext(coinApiProvider)
	case "coinapi":
		genesisProvider = coinApiProvider
		coinApiProvider.SetNext(binanceProvider)
		binanceProvider.SetNext(coinbaseProvider)
	default:
		genesisProvider = binanceProvider
		binanceProvider.SetNext(coinbaseProvider)
		coinbaseProvider.SetNext(coinApiProvider)
	}
	return &genesisProvider
}
