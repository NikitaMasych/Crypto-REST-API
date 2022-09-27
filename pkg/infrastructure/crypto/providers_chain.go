package crypto

import (
	"GenesisTask/pkg/application"
)

func NewProvidersChain() *application.ProvidersChain {
	coinbaseProvider := NewCoinbaseProvider()
	binanceProvider := NewBinanceProvider()
	coinApiProvider := NewCoinApiProvider()

	coinbaseProvider.SetNext(&binanceProvider)
	binanceProvider.SetNext(&coinApiProvider)

	return &coinbaseProvider
}
