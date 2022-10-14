package models

import (
	"producer/config"
)

type Subscription struct {
	EmailAddress
	CurrencyPair
}

func NewSubscription(e EmailAddress, p CurrencyPair) *Subscription {
	return &Subscription{e, p}
}

func (s *Subscription) GetEmailAddress() *EmailAddress {
	return &s.EmailAddress
}

func (s *Subscription) GetCurrencyPair() *CurrencyPair {
	return &s.CurrencyPair
}

func (s *Subscription) ToString() string {
	return s.GetEmailAddress().ToString() + config.EmailAddressSeparator +
		s.GetCurrencyPair().ToString()
}
