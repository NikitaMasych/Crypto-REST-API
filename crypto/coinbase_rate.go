package crypto

import (
	"GenesisTask/config"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type CoinbaseProvider struct {
	Response struct {
		BaseCurrency   string `json:"base"`
		QuotedCurrency string `json:"currency"`
		Price          string `json:"amount"`
	} `json:"data"`
}

func (p *CoinbaseProvider) GetConfigCurrencyRate() (float64, error) {
	cfg := config.Get()
	CryptoApiUrl := fmt.Sprintf(
		cfg.CoinbaseApiFormatUrl, cfg.BaseCurrency, cfg.QuotedCurrency)

	resp, err := http.Get(CryptoApiUrl)
	if err != nil {
		return 0, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return 0, err
	}

	return strconv.ParseFloat(p.Response.Price, 64)
}

type CoinbaseProviderCreator struct{}

func (p *CoinbaseProviderCreator) CreateProvider() CryptoProvider {
	return new(CoinbaseProvider)
}
