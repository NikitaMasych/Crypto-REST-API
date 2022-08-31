package crypto

import (
	"GenesisTask/config"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type decodedResponse struct {
	Data struct {
		BaseCurrency   string `json:"base"`
		QuotedCurrency string `json:"currency"`
		Price          string `json:"amount"`
	} `json:"data"`
}

func GetConfigCurrencyRate() (float64, error) {
	cfg := config.Get()
	CryptoApiUrl := fmt.Sprintf(
		cfg.CryptoApiFormatUrl, cfg.BaseCurrency, cfg.QuotedCurrency)

	resp, err := http.Get(CryptoApiUrl)
	if err != nil {
		return 0, err
	}

	var cryptoRate decodedResponse
	if err := json.NewDecoder(resp.Body).Decode(&cryptoRate); err != nil {
		return 0, err
	}
	return strconv.ParseFloat(cryptoRate.Data.Price, 64)
}
