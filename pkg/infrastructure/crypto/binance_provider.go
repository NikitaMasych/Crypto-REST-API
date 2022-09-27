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

type BinanceProvider struct {
	Response struct {
		Price string `json:"price"`
	}
	next *application.ProvidersChain
}

func NewBinanceProvider() application.ProvidersChain {
	return &BinanceProvider{}
}

func (p *BinanceProvider) SetNext(next *application.ProvidersChain) {
	p.next = next
}

func (p *BinanceProvider) GetRate(pair models.CurrencyPair) (models.CurrencyRate, error) {
	rate, err := p.getRate(pair)
	if err != nil {
		if p.next == nil {
			return rate, err
		}
		return (*p.next).GetRate(pair)
	}
	return rate, err
}

func (p *BinanceProvider) getRate(pair models.CurrencyPair) (models.CurrencyRate, error) {
	BinanceApiUrl := fmt.Sprintf(
		config.Get().BinanceApiFormatUrl, pair.GetBase(), pair.GetQuote())

	resp, err := resty.R().Get(BinanceApiUrl)
	if err != nil {
		return *models.NewCurrencyRate(pair, -1), err
	}

	go logger.LogProviderResponse("Binance", resp)

	if err := json.Unmarshal(resp.Body, &p.Response); err != nil {
		return *models.NewCurrencyRate(pair, -1), err
	}
	rate, err := strconv.ParseFloat(p.Response.Price, 64)
	if err != nil {
		return *models.NewCurrencyRate(pair, -1), err
	}

	return *models.NewCurrencyRate(pair, rate), err
}
