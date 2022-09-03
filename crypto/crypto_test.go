package crypto

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestIsCryptoRateValid(t *testing.T) {
	err := godotenv.Load("./../.env")
	if err != nil {
		log.Fatal(err)
	}

	_, err = GetConfigCurrencyRate()
	if err != nil {
		t.Error(err)
	}
}
