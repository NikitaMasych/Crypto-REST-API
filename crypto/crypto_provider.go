package crypto

import (
	"GenesisTask/config"
	"log"
)

type CryptoProvider interface {
	getCurrencyRate(base, quoted string) (float64, error)
}

type CryptoProviderCreator interface {
	createProvider() CryptoProvider
}

func GetProviderCurrencyRate(p CryptoProviderCreator) (float64, error) {
	cfg := config.Get()
	return p.createProvider().getCurrencyRate(cfg.BaseCurrency, cfg.QuotedCurrency)
}

func NewConfigProviderCreator() CryptoProviderCreator {
	provider := config.Get().CryptoCurrencyProvider
	switch provider {
	case "binance":
		return new(BinanceProviderCreator)
	case "coinbase":
		return new(CoinbaseProviderCreator)
	case "coinapi":
		return new(CoinApiProviderCreator)
	default:
		log.Fatal("Unknown provider")
	}
	// never reach here, golang requirement
	return *new(CryptoProviderCreator)
}
