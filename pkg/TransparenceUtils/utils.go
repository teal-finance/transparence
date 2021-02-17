package transparenceutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"net/http"
	"os"

	"transparence/pkg/erc20adapter"
)

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func ParseDecimals(totalSupply *big.Int, decimals uint8) *big.Float {
	tmp := new(big.Float)
	tmp.SetString(totalSupply.String())
	value := new(big.Float).Quo(tmp, big.NewFloat(math.Pow10(int(decimals))))
	return value
}

func HttpRequest(address string) []byte {
	resp, err := http.Get(address)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}

func PercentOf(part float64, total float64) float64 {
	return (part * 100) / total
}

func Verif(supplyOnReserve *big.Float, supplyOnToken *big.Float) bool {
	fsupplyOnReserve, _ := supplyOnReserve.Float64()
	fsupplyOnToken, _ := supplyOnToken.Float64()
	if fsupplyOnReserve < fsupplyOnToken {
		return false
	} else {
		return true
	}
}

func ReadConfigFile(configFile string) erc20adapter.PeggedTokens {
	file, _ := os.Open(configFile)
	defer file.Close()
	decoder := json.NewDecoder(file)
	tokens := erc20adapter.PeggedTokens{}
	err := decoder.Decode(&tokens)
	if err != nil {
		fmt.Println("error reading conf file :", err)
	}
	return tokens
}
