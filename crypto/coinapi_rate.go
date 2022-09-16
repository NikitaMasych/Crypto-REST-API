package crypto

import (
	"GenesisTask/config"
	"GenesisTask/logger"
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v0"
)

type CoinApiProvider struct {
	Response struct {
		Price float64 `json:"rate"`
	}
}

func (p *CoinApiProvider) getCurrencyRate(base, quoted string) (float64, error) {
	cfg := config.Get()
	CoinApiURL := fmt.Sprintf(
		cfg.CoinApiFormatURL, base, quoted)

	resp, err := resty.R().SetHeader("X-CoinAPI-Key", cfg.CoinApiKey).
		Get(CoinApiURL)
	if err != nil {
		return 0, err
	}

	go logger.LogProviderResponse("CoinApi", resp)

	if err := json.Unmarshal(resp.Body, &p.Response); err != nil {
		return 0, err
	}
	return p.Response.Price, nil
}

type CoinApiProviderCreator struct{}

func (p *CoinApiProviderCreator) createProvider() CryptoProvider {
	return new(CoinApiProvider)
}
