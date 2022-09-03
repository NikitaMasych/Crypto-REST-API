package crypto

import "testing"

func TestThatCryptoRateReceived(t *testing.T) {

	_, err := GetConfigCurrencyRate()

	if err != nil {
		t.Error(err)
	}
}
