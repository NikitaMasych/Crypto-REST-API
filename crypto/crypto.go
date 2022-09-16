package crypto

import (
	"GenesisTask/cache"
	"GenesisTask/config"
	"log"
)

type CryptoProvider interface {
	GetConfigCurrencyRate() (float64, error)
}

type CryptoProviderCreator interface {
	CreateProvider() CryptoProvider
}

func GetProviderCryptoRate(p CryptoProviderCreator) (float64, error) {
	return p.CreateProvider().GetConfigCurrencyRate()
}

func EnvProviderDescriptor() CryptoProviderCreator {
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

func GetCryptoRate() (float64, error) {
	price, err := new(BinanceProviderCreator).CreateProvider().GetConfigCurrencyRate()
	if err != nil {
		log.Print("Can't get rate from coinbase")
		price, err = new(BinanceProviderCreator).CreateProvider().GetConfigCurrencyRate()
		if err != nil {
			log.Print("Can't get rate from binance")
			price, err = new(CoinApiProviderCreator).CreateProvider().GetConfigCurrencyRate()
			if err != nil {
				log.Print("Can't get rate from coinapi")
				return 0, err
			}
		}
	}
	if err == nil {
		cache.AddCryptoRateToCache(price)
	}
	return price, err
}
