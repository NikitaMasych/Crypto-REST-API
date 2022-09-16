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

func (p *BinanceProvider) getCurrencyRate(base, quoted string) (float64, error) {
	BinanceApiUrl := fmt.Sprintf(
		config.Get().BinanceApiFormatUrl, base, quoted)

	resp, err := resty.R().Get(BinanceApiUrl)
	if err != nil {
		return 0, err
	}

	go logger.LogProviderResponse("Binance", resp)

	if err := json.Unmarshal(resp.Body, &p.Response); err != nil {
		return 0, err
	}
	return strconv.ParseFloat(p.Response.Price, 64)
}

type BinanceProviderCreator struct{}

func (p *BinanceProviderCreator) createProvider() CryptoProvider {
	return new(BinanceProvider)
}
