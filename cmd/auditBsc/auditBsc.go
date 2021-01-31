package main

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	transparenceutils "transparence/pkg/TransparenceUtils"
	"transparence/pkg/erc20adapter"
)

const CONFIG_FILE = "binancePeggedTokens.json"
const INFURA_KEY = ""
const IP_API_INFURA = "Https://mainnet.infura.io/v3/" + INFURA_KEY
const IP_API_BINANCESMARTCHAIN = "wss://bsc-ws-node.nariox.org:443"
const IP_API_BINANCECHAIN = "Https://dex.binance.org/api/v1/tokens"

func main() {
	auditBinanceTokens(IP_API_INFURA)
}

func auditBinanceTokens(ipAddress string) {
	total := big.NewFloat(0)
	var name, symbol string

	tokens := transparenceutils.ReadConfigFile(CONFIG_FILE)
	client := erc20adapter.NewClientConnection(ipAddress)

	for i := 0; i < len(tokens.PeggedTokens); i++ {
		tokenAddress := common.HexToAddress(tokens.PeggedTokens[i].CONTRACT_ADDRESS)
		instance := erc20adapter.NewToken(tokenAddress, client)

		if tokens.PeggedTokens[i].SYMBOL == "MKR" {
			name = "Maker"
			symbol = tokens.PeggedTokens[i].SYMBOL
		} else {
			name = erc20adapter.GetName(instance)
			symbol = erc20adapter.GetSymbol(instance)
		}

		decimals := erc20adapter.GetDecimals(instance)

		fmt.Printf("\n---- Auditing Binance Pegged Token : '%s'  (%s)\n", name, symbol)
		fmt.Printf("- Ethereum reserve addresses : %s \n", tokens.PeggedTokens[i].CONTRACT_ADDRESS)
		fmt.Printf("- Binance Chain reserve addresses : %s \n", tokens.PeggedTokens[i].BNC_PEGGED_ADDRESSES)
		fmt.Printf("- Binance Smart Chain reserve addresses : %s \n", tokens.PeggedTokens[i].BSC_PEGGED_ADDRESSES)

		totalReserve := erc20adapter.GetBalanceOfToken(tokens.PeggedTokens[i].RESERVE_ADDRESS, instance)
		ftotalReserve := transparenceutils.ParseDecimals(totalReserve, decimals)
		total.Add(total, ftotalReserve)
		executeAuditOnEth(tokens.PeggedTokens[i], ftotalReserve)
	}
}

func executeAuditOnEth(tokenAudited erc20adapter.PeggedToken, supplyOnEth *big.Float) {
	supplyOnBitcoin := big.NewFloat(0)
	tokenSupply := new(big.Float)

	binanceTokens, _ := getBinanceChainTokens()
	for i := 0; i < len(tokenAudited.BNC_PEGGED_ADDRESSES); i++ {
		for _, item := range binanceTokens {
			if item.ContractAddress == tokenAudited.BNC_PEGGED_ADDRESSES[i] {
				tokenSupply.SetString(item.TotalSupply)
				supplyOnBitcoin.Add(supplyOnBitcoin, tokenSupply)
			}
		}
	}

	client := erc20adapter.NewClientConnection(IP_API_BINANCESMARTCHAIN)

	for i := 0; i < len(tokenAudited.BSC_PEGGED_ADDRESSES); i++ {
		tokenAddress := common.HexToAddress(tokenAudited.BSC_PEGGED_ADDRESSES[i])
		instance := erc20adapter.NewToken(tokenAddress, client)
		decimals := erc20adapter.GetDecimals(instance)
		totalSupply := erc20adapter.GetTotalSupply(instance)
		ftotalSupply := transparenceutils.ParseDecimals(totalSupply, decimals)
		supplyOnBitcoin.Add(supplyOnBitcoin, ftotalSupply)
	}
	auditResult := transparenceutils.Verif(supplyOnEth, supplyOnBitcoin)
	fmt.Printf("- Reserve on Ethereum : %f \n- TotalSupply on BNC+BSC: %f\n---- Audit result : %s \n", supplyOnEth, supplyOnBitcoin, auditResult)
}

func getBinanceChainTokens() ([]erc20adapter.BEP2Token, error) {
	var tokens []erc20adapter.BEP2Token
	body := transparenceutils.HttpRequest(IP_API_BINANCECHAIN)
	if err := json.Unmarshal(body, &tokens); err != nil {
		return nil, err
	}
	return tokens, nil
}
