package erc20adapter

import (
	"log"
	"math/big"

	"transparence/pkg/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Types used to represent the tokens of the JSON config file
type PeggedToken struct {
	BLOCKCHAIN_PLATFORM  string   `json:"BLOCKCHAIN_PLATFORM"`
	SYMBOL               string   `json:"SYMBOL"`
	CONTRACT_ADDRESS     string   `json:"CONTRACT_ADDRESS"`
	RESERVE_ADDRESS      string   `json:"RESERVE_ADDRESS,omitempty"`
	BTC_ADDRESSES        []string `json:"BTC_ADDRESSES,omitempty"`
	BNC_PEGGED_ADDRESSES []string `json:"BNC_PEGGED_ADDRESSES,omitempty"`
	BSC_PEGGED_ADDRESSES []string `json:"BSC_PEGGED_ADDRESSES,omitempty"`
}

type BEP2Token struct {
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	OrigSymbol       string `json:"original_symbol"`
	TotalSupply      string `json:"total_supply"`
	Owner            string `json:"owner"`
	Mintable         bool   `json:"mintable"`
	ContractAddress  string `json:"contract_address,omitempty"`
	ContractDecimals int8   `json:"contract_decimals,omitempty"`
}

type PeggedTokens struct {
	PeggedTokens []PeggedToken `json:"tokens"`
}

func NewClientConnection(ipAddress string) *ethclient.Client {
	client, err := ethclient.Dial(ipAddress)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func NewToken(tokenAddress common.Address, client *ethclient.Client) *token.Token {
	tokenInstance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	return tokenInstance
}

func GetName(tokenERC20 *token.Token) string {
	name, err := tokenERC20.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return name
}

func GetSymbol(tokenERC20 *token.Token) string {
	symbol, err := tokenERC20.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return symbol
}

func GetDecimals(tokenERC20 *token.Token) uint8 {
	decimals, err := tokenERC20.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return decimals
}

func GetTotalSupply(tokenERC20 *token.Token) *big.Int {
	totalSupply, err := tokenERC20.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return totalSupply
}

func GetBalanceOfToken(addressToQuery string, tokenERC20 *token.Token) *big.Int {
	address := common.HexToAddress(addressToQuery)
	bal, err := tokenERC20.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	return bal
}
