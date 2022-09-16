package crypto

import (
	"GenesisTask/config"
	"GenesisTask/logger"
	"encoding/json"
	"fmt"
	"strconv"

	"gopkg.in/resty.v0"
)

type BinanceProvider struct {
	Response struct {
		Price string `json:"price"`
	}
}

func (p *BinanceProvider) GetConfigCurrencyRate() (float64, error) {
	cfg := config.Get()
	BinanceApiUrl := fmt.Sprintf(
		cfg.BinanceApiFormatUrl, cfg.BaseCurrency, cfg.QuotedCurrency)

	resp, err := resty.R().Get(BinanceApiUrl)
	if err != nil {
		return 0, err
	}

	go logger.AddProviderResponseToLog("Binance", resp)

	if err := json.Unmarshal(resp.Body, &p.Response); err != nil {
		return 0, err
	}
	return strconv.ParseFloat(p.Response.Price, 64)
}

type BinanceProviderCreator struct{}

func (p *BinanceProviderCreator) CreateProvider() CryptoProvider {
	return new(BinanceProvider)
}
