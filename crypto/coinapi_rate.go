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
		Time           string  `json:"time"`
		BaseCurrency   string  `json:"asset_id_base"`
		QuotedCurrency string  `json:"asset_id_quote"`
		Price          float64 `json:"rate"`
	}
}

func (p *CoinApiProvider) GetConfigCurrencyRate() (float64, error) {
	cfg := config.Get()
	CoinApiURL := fmt.Sprintf(
		cfg.CoinApiFormatURL, cfg.BaseCurrency, cfg.QuotedCurrency)

	resp, err := resty.R().SetHeader("X-CoinAPI-Key", cfg.CoinApiKey).
		Get(CoinApiURL)
	if err != nil {
		return 0, err
	}
	logger.AddProviderResponseToLog(resp.RawResponse)
	if err := json.Unmarshal(resp.Body, &p.Response); err != nil {
		return 0, err
	}
	return p.Response.Price, nil
}

type CoinApiProviderCreator struct{}

func (p *CoinApiProviderCreator) CreateProvider() CryptoProvider {
	return new(CoinApiProvider)
}
