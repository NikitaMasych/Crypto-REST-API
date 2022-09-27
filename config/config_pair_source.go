package config

import (
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/domain/models"
)

type ConfigPairSource struct{}

func NewConfigPairSource() application.PairSource {
	return &ConfigPairSource{}
}

func (s *ConfigPairSource) GetPair() models.CurrencyPair {
	return *models.NewCurrencyPair(Get().BaseCurrency, Get().QuoteCurrency)
}
