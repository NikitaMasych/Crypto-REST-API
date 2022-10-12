package crypto

import (
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/domain/models"

	"encoding/json"
	"fmt"
	"strconv"

	"gopkg.in/resty.v0"
)

type BinanceProvider struct {
	Response struct {
		Price string `json:"price"`
	}
	next *application.ProvidersChain
}

func NewBinanceProvider() application.ProvidersChain {
	return &BinanceProvider{}
}

func (p *BinanceProvider) SetNext(next application.ProvidersChain) {
	p.next = &next
}

func (p *BinanceProvider) GetRate(pair models.CurrencyPair) (models.CurrencyRate, error) {
	rate, err := p.getRate(pair)
	if err != nil {
		if p.next == nil {
			logger.LogError(err)
			return rate, err
		}
		logger.LogDebug("Calling next provider in the chain")
		return (*p.next).GetRate(pair)
	}
	return rate, err
}

func (p *BinanceProvider) getRate(pair models.CurrencyPair) (models.CurrencyRate, error) {
	BinanceApiUrl := fmt.Sprintf(
		BinanceApiFormatUrl, pair.GetBase(), pair.GetQuote())

	resp, err := resty.R().Get(BinanceApiUrl)
	logger.LogDebug("Requested from " + BinanceApiUrl)
	timestamp := resp.ReceivedAt
	if err != nil {
		return *models.NewCurrencyRate(pair, -1, timestamp), err
	}

	go logger.LogDebug(ComposeProviderResponseLog(timestamp, "Binance", resp))

	if err := json.Unmarshal(resp.Body, &p.Response); err != nil {
		return *models.NewCurrencyRate(pair, -1, timestamp), err
	}
	rate, err := strconv.ParseFloat(p.Response.Price, 64)
	if err != nil {
		return *models.NewCurrencyRate(pair, -1, timestamp), err
	}

	return *models.NewCurrencyRate(pair, rate, timestamp), err
}
