package crypto

import (
	"GenesisTask/config"
	"log"
)

type CryptoProvider interface {
	GetConfigCurrencyRate() (float64, error)
}

type CryptoProviderCreator interface {
	CreateProvider() CryptoProvider
}

func GetCryptoRate(p CryptoProviderCreator) (float64, error) {
	return p.CreateProvider().GetConfigCurrencyRate()
}

func EnvProviderDescriptor() CryptoProviderCreator {
	provider := config.Get().CryptoCurrencyProvider
	switch provider {
	case "binance":
		return new(BinanceProviderCreator)
	case "coinbase":
		return new(CoinbaseProviderCreator)
	default:
		log.Fatal("Unknown provider")
	}
	// never reach here, golang requirement
	return new(CoinbaseProviderCreator)
}
