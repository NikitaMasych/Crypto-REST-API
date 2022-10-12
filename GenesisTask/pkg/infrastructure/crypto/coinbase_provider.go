package crypto

import (
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/domain/models"
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

func (p *CoinbaseProvider) SetNext(next application.ProvidersChain) {
	p.next = &next
}

func (p *CoinbaseProvider) GetRate(pair models.CurrencyPair) (models.CurrencyRate, error) {
	rate, err := p.getRate(pair)
	if err != nil {
		if p.next == nil {
			return rate, err
		}
		logger.LogDebug("Calling next provider in the chain")
		return (*p.next).GetRate(pair)
	}
	return rate, err
}

func (p *CoinbaseProvider) getRate(pair models.CurrencyPair) (models.CurrencyRate, error) {
	CoinbaseApiUrl := fmt.Sprintf(
		CoinbaseApiFormatUrl, pair.GetBase(), pair.GetQuote())
	resp, err := resty.R().Get(CoinbaseApiUrl)
	logger.LogDebug("Requested from " + CoinbaseApiUrl)
	timestamp := resp.ReceivedAt
	if err != nil {
		return *models.NewCurrencyRate(pair, -1, timestamp), err
	}

	go logger.LogDebug(ComposeProviderResponseLog(timestamp, "Coinbase", resp))

	if err := json.Unmarshal(resp.Body, &p.Response); err != nil {
		return *models.NewCurrencyRate(pair, -1, timestamp), err
	}
	rate, err := strconv.ParseFloat(p.Response.Data.Price, 64)
	if err != nil {
		return *models.NewCurrencyRate(pair, -1, timestamp), err
	}

	return *models.NewCurrencyRate(pair, rate, timestamp), err
}
