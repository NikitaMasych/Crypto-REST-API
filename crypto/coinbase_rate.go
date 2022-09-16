package crypto

import (
	"GenesisTask/config"
	"GenesisTask/logger"
	"encoding/json"
	"fmt"
	"strconv"

	"gopkg.in/resty.v0"
)

type CoinbaseProvider struct {
	Response struct {
		Price string `json:"amount"`
	} `json:"data"`
}

func (p *CoinbaseProvider) getCurrencyRate(base, quoted string) (float64, error) {
	CoinbaseApiUrl := fmt.Sprintf(
		config.Get().CoinbaseApiFormatUrl, base, quoted)

	resp, err := resty.R().Get(CoinbaseApiUrl)
	if err != nil {
		return 0, err
	}

	go logger.LogProviderResponse("Coinbase", resp)

	if err := json.Unmarshal(resp.Body, &p); err != nil {
		return 0, err
	}
	return strconv.ParseFloat(p.Response.Price, 64)
}

type CoinbaseProviderCreator struct{}

func (p *CoinbaseProviderCreator) createProvider() CryptoProvider {
	return new(CoinbaseProvider)
}
