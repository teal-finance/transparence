package coingecko

import (
	"encoding/json"
	"log"
	transparenceutils "transparence/pkg/TransparenceUtils"
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
	body := transparenceutils.HttpRequest(IP_API_COINGECKO)
	data := CoinsMarket{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}

	return data[0]
}
