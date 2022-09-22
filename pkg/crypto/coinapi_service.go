package crypto

type coinApiService struct {
	next *Cryptochain
}

func NewCoinApiService() Cryptochain {
	return &coinApiService{}
}

func (b *coinApiService) SetNext(next *Cryptochain) {
	b.next = next
}

func (b *coinApiService) GetCurrencyRate(base, quoted string) (float64, error) {
	rate, err := new(CoinApiProvider).getCurrencyRate(base, quoted)
	if err != nil {
		if b.next == nil {
			return 0, err
		}
		rate, err = (*b.next).GetCurrencyRate(base, quoted)
	}
	return rate, err
}
