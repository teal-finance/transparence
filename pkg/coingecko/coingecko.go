package coingecko

import (
	"encoding/json"
	"log"

	"transparence/pkg/http"
)

const IP_API_COINGECKO = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=bitcoin"

// Types used to represent Coingecko reponses (a lot of more infos are also available but not stored)
type CoinsMarketItem struct {
	Symbol            string  `json:"symbol"`
	Name              string  `json:"name"`
	TotalSupply       float64 `json:"total_supply"`
	CirculatingSupply float64 `json:"circulating_supply"`
	CurrentPrice      float64 `json:"current_price"`
	MarketCap         float64 `json:"market_cap"`
	MarketCapRank     int16   `json:"market_cap_rank"`
}

type CoinsMarket []CoinsMarketItem

func QueryBitcoinInfosFromGecko() CoinsMarketItem {
	body, err := http.Get(IP_API_COINGECKO)
	if err != nil {
		log.Fatal(err)
	}
	data := CoinsMarket{}
	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatalln(err)
	}

	return data[0]
}
