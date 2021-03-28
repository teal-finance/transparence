package math

import (
	"math"
	"math/big"
)

func Percent(part float64, total float64) float64 {
	return (part * 100) / total
}

func ParseDecimals(totalSupply *big.Int, decimals uint8) *big.Float {
	tmp := new(big.Float)
	tmp.SetString(totalSupply.String())
	value := new(big.Float).Quo(tmp, big.NewFloat(math.Pow10(int(decimals))))
	return value
}

func Verif(supplyOnReserve *big.Float, supplyOnToken *big.Float) string {
	fsupplyOnReserve, _ := supplyOnReserve.Float64()
	fsupplyOnToken, _ := supplyOnToken.Float64()
	if fsupplyOnReserve < fsupplyOnToken {
		return "KO"
	} else {
		return "OK"
	}
}
