package BitcoinPrice

import (
	"encoding/json"
	"net/http"
)

const coinbaseAPI = "https://api.coinbase.com/v2/prices/BTC-UAH/spot"

type data struct {
	Price string `json:"amount"`
}

type cryptoValue struct {
	Data data `json:"data"`
}

func GetBitcoinPrice() (string, int) {

	code := 200 // Initially assume that everything will be file :)

	response, err := http.Get(coinbaseAPI)
	if err != nil {
		code = 400
		return "", code
	}
	var decodedResponse cryptoValue
	if err := json.NewDecoder(response.Body).Decode(&decodedResponse); err != nil {
		// Invalid JSON
		code = 400
		return "", code
	}

	return decodedResponse.Data.Price, code
}
