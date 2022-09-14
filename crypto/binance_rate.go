package crypto

import (
	"GenesisTask/config"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

	resp, err := http.Get(BinanceApiUrl)
	if err != nil {
		return 0, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&p.Response); err != nil {
		return 0, err
	}
	return strconv.ParseFloat(p.Response.Price, 64)
}

type BinanceProviderCreator struct{}

func (p *BinanceProviderCreator) CreateProvider() CryptoProvider {
	return new(BinanceProvider)
}
