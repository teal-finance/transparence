package cTokenAdapter

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"transparence/pkg/comptroller"
	"transparence/pkg/crToken"
)

func NewClientConnection(ipAddress string) *ethclient.Client {
	client, err := ethclient.Dial(ipAddress)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func NewCrToken(tokenAddress common.Address, client *ethclient.Client) *crToken.CrToken {
	tokenInstance, err := crToken.NewCrToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	return tokenInstance
}

func NewComptroller(tokenAddress common.Address, client *ethclient.Client) *comptroller.Comptroller {
	instance, err := comptroller.NewComptroller(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	return instance
}
func GetName(cTokenInstance *crToken.CrToken) string {
	name, err := cTokenInstance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return name
}

func GetSymbol(cTokenInstance *crToken.CrToken) string {
	symbol, err := cTokenInstance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return symbol
}

func GetDecimals(cTokenInstance *crToken.CrToken) uint8 {
	decimals, err := cTokenInstance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return decimals
}

func GetTotalSupply(cTokenInstance *crToken.CrToken) *big.Int {
	totalSupply, err := cTokenInstance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return totalSupply
}

func GetBalanceOfToken(addressToQuery string, cTokenInstance *crToken.CrToken) *big.Int {
	address := common.HexToAddress(addressToQuery)
	bal, err := cTokenInstance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	return bal
}

func GetTotalBorrows(cTokenInstance *crToken.CrToken) *big.Int {
	totalBorrows, err := cTokenInstance.TotalBorrows(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return totalBorrows
}

func GetTotalReserves(cTokenInstance *crToken.CrToken) *big.Int {
	totalReserves, err := cTokenInstance.TotalReserves(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return totalReserves
}

func GetCash(cTokenInstance *crToken.CrToken) *big.Int {
	cash, err := cTokenInstance.GetCash(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return cash
}

func GetUnderlying(cTokenInstance *crToken.CrToken) common.Address {
	underlyingAddress, err := cTokenInstance.Underlying(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return underlyingAddress
}

func GetCompAddress(currentComptroller *comptroller.Comptroller) common.Address {
	platformTokenAddress, err := currentComptroller.GetCompAddress(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return platformTokenAddress
}

func GetAllMarkets(currentComptroller *comptroller.Comptroller) []common.Address {
	markets, err := currentComptroller.GetAllMarkets(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return markets
}
