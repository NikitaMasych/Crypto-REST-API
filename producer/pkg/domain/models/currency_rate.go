package models

import (
	"time"
)

type CurrencyRate struct {
	CurrencyPair
	price     float64
	timestamp time.Time
}

func NewCurrencyRate(p CurrencyPair, price float64, timestamp time.Time) *CurrencyRate {
	return &CurrencyRate{p, price, timestamp}
}

func (r *CurrencyRate) GetCurrencyPair() *CurrencyPair {
	return &r.CurrencyPair
}

func (r *CurrencyRate) GetPrice() float64 {
	return r.price
}

func (r *CurrencyRate) GetTimestamp() time.Time {
	return r.timestamp
}
