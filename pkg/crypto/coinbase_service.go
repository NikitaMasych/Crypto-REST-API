package crypto

type coinbaseService struct {
	next *Cryptochain
}

func NewCoinbaseService() Cryptochain {
	return &coinbaseService{}
}

func (b *coinbaseService) SetNext(next *Cryptochain) {
	b.next = next
}

func (b *coinbaseService) GetCurrencyRate(base, quoted string) (float64, error) {
	rate, err := new(CoinbaseProvider).getCurrencyRate(base, quoted)
	if err != nil {
		if b.next == nil {
			return 0, err
		}
		rate, err = (*b.next).GetCurrencyRate(base, quoted)
	}
	return rate, err
}
