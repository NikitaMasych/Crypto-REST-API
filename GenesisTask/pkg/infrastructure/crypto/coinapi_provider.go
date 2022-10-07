package crypto

import (
	"GenesisTask/config"
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/domain/models"
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

func (p *CoinApiProvider) SetNext(next application.ProvidersChain) {
	p.next = &next
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
	CoinApiUrl := fmt.Sprintf(
		CoinApiFormatUrl, pair.GetBase(), pair.GetQuote())

	resp, err := resty.R().SetHeader("X-CoinAPI-Key", config.CoinApiKey).Get(CoinApiUrl)
	timestamp := resp.ReceivedAt
	if err != nil {
		return *models.NewCurrencyRate(pair, -1, timestamp), err
	}

	go logger.LogInfo(ComposeProviderResponseLog(timestamp, "CoinApi", resp))

	if err := json.Unmarshal(resp.Body, &p.Response); err != nil {
		return *models.NewCurrencyRate(pair, -1, timestamp), err
	}
	rate := p.Response.Price

	return *models.NewCurrencyRate(pair, rate, timestamp), err
}
