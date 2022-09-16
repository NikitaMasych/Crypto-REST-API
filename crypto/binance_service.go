package crypto

type binanceService struct {
	next *Cryptochain
}

func NewBinanceService() Cryptochain {
	return &binanceService{}
}

func (b *binanceService) SetNext(next *Cryptochain) {
	b.next = next
}

func (b *binanceService) GetCurrencyRate(base, quoted string) (float64, error) {
	rate, err := new(BinanceProvider).getCurrencyRate(base, quoted)
	if err != nil {
		if b.next == nil {
			return 0, err
		}
		rate, err = (*b.next).GetCurrencyRate(base, quoted)
	}
	return rate, err
}
