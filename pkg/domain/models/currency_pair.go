package models

type CurrencyPair struct {
	base  string
	quote string
}

func NewCurrencyPair(base string, quote string) *CurrencyPair {
	return &CurrencyPair{base, quote}
}

func (p CurrencyPair) GetBase() string {
	return p.base
}

func (p CurrencyPair) GetQuote() string {
	return p.quote
}
