package crypto

import (
	"GenesisTask/config"
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/domain/models"
	"GenesisTask/pkg/infrastructure/logger"
	"encoding/json"
	"fmt"
	"strconv"

	"gopkg.in/resty.v0"
)

type CoinbaseProvider struct {
	Response struct {
		Data struct {
			Price string `json:"amount"`
		} `json:"data"`
	}
	next *application.ProvidersChain
}

func NewCoinbaseProvider() application.ProvidersChain {
	return &CoinbaseProvider{}
}

func (p *CoinbaseProvider) SetNext(next *application.ProvidersChain) {
	p.next = next
}

func (p *CoinbaseProvider) GetRate(pair models.CurrencyPair) (models.CurrencyRate, error) {
	rate, err := p.getRate(pair)
	if err != nil {
		if p.next == nil {
			return rate, err
		}
		return (*p.next).GetRate(pair)
	}
	return rate, err
}

func (p *CoinbaseProvider) getRate(pair models.CurrencyPair) (models.CurrencyRate, error) {
	CoinbaseApiUrl := fmt.Sprintf(
		config.Get().CoinbaseApiFormatUrl, pair.GetBase(), pair.GetQuote())

	resp, err := resty.R().Get(CoinbaseApiUrl)
	if err != nil {
		return *models.NewCurrencyRate(pair, -1), err
	}

	go logger.LogProviderResponse("Coinbase", resp)

	if err := json.Unmarshal(resp.Body, &p.Response); err != nil {
		return *models.NewCurrencyRate(pair, -1), err
	}
	rate, err := strconv.ParseFloat(p.Response.Data.Price, 64)
	if err != nil {
		return *models.NewCurrencyRate(pair, -1), err
	}

	return *models.NewCurrencyRate(pair, rate), err
}
