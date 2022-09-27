package models

type CurrencyRate struct {
	CurrencyPair
	price float64
}

func NewCurrencyRate(p CurrencyPair, price float64) *CurrencyRate {
	return &CurrencyRate{p, price}
}

func (r *CurrencyRate) GetCurrencyPair() *CurrencyPair {
	return &r.CurrencyPair
}

func (r *CurrencyRate) GetPrice() float64 {
	return r.price
}
