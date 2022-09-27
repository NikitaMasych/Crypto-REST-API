package crypto

import (
	"GenesisTask/config"
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/domain/models"
	"GenesisTask/pkg/infrastructure/logger"
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v0"
)

type CoinApiProvider struct {
	Response struct {
		Price float64 `json:"rate"`
	}
	next *application.ProvidersChain
}

func NewCoinApiProvider() application.ProvidersChain {
	return &CoinApiProvider{}
}

func (p *CoinApiProvider) SetNext(next *application.ProvidersChain) {
	p.next = next
}

func (p *CoinApiProvider) GetRate(pair models.CurrencyPair) (models.CurrencyRate, error) {
	rate, err := p.getRate(pair)
	if err != nil {
		if p.next == nil {
			return rate, err
		}
		return (*p.next).GetRate(pair)
	}
	return rate, err
}

func (p *CoinApiProvider) getRate(pair models.CurrencyPair) (models.CurrencyRate, error) {
	cfg := config.Get()
	CoinApiUrl := fmt.Sprintf(
		cfg.CoinApiFormatURL, pair.GetBase(), pair.GetQuote())

	resp, err := resty.R().SetHeader("X-CoinAPI-Key", cfg.CoinApiKey).Get(CoinApiUrl)
	if err != nil {
		return *models.NewCurrencyRate(pair, -1), err
	}

	go logger.LogProviderResponse("CoinApi", resp)

	if err := json.Unmarshal(resp.Body, &p.Response); err != nil {
		return *models.NewCurrencyRate(pair, -1), err
	}
	rate := p.Response.Price

	return *models.NewCurrencyRate(pair, rate), err
}
